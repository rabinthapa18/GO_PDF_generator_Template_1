package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"grrow_pdf/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	npdf "github.com/dslipak/pdf"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
)

func AddToTemplate(template string, definition models.Definitions, value models.Values) []byte {

	// download template
	temp := downloadFiles(template)
	var seal, logo []byte
	ioutil.WriteFile("temp.pdf", temp, 0644)

	// download seal and logo from value model
	for _, v := range value.Images {
		if v.Name == "seal" {
			seal = downloadFiles(v.ObjectKey)
		}
		if v.Name == "logo" {
			logo = downloadFiles(v.ObjectKey)
		}
	}

	// create new pdf

	pdf := gofpdf.New("P", "mm", "A4", "")

	// get number of pages from template file ============================
	contents, err := npdf.Open("temp.pdf")
	if err != nil {
		panic(err)
	}
	numberOfPages := contents.NumPage()

	// changing template received from api to readseeker==================
	templateFile := io.ReadSeeker(bytes.NewReader(temp))

	// Add a page to the document ========================================
	for i := 1; i <= numberOfPages; i++ {
		page := gofpdi.ImportPageFromStream(pdf, &templateFile, i, "/MediaBox")

		pdf.AddPage()

		// Draw imported template onto page
		gofpdi.UseImportedTemplate(pdf, page, 0, 0, 215, 0)
	}

	final := pdf.PageNo()
	println(final)

	// write data to pdf
	writeDataToFile(value, definition, pdf, seal, logo)

	pdf.SetPage(final)

	// save pdf to file
	err = pdf.OutputFileAndClose("temp.pdf")
	if err != nil {
		println("error during saving pdf")
	}

	// change pdf to bytes
	pdfBytes, err := ioutil.ReadFile("temp.pdf")
	if err != nil {
		panic(err)
	}

	//delete temp file
	// err = os.Remove("temp.pdf")
	// if err != nil {
	// 	panic(err)
	// }

	// delete files from s3
	deleteFile(template)
	for _, v := range value.Images {
		if v.Name == "seal" {
			deleteFile(v.ObjectKey)
		}
		if v.Name == "logo" {
			deleteFile(v.ObjectKey)
		}
	}

	return pdfBytes

}

func downloadFiles(fileName string) []byte {
	svc := GetS3()
	pdfTemplate := &s3.GetObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String(fileName),
	}
	req, out := svc.GetObject(context.TODO(), pdfTemplate)
	if out != nil {
		fmt.Println(out.Error())
	}
	// reading the template file received via API
	file, _ := ioutil.ReadAll(req.Body)

	req.Body.Close()

	return file
}

func writeDataToFile(value models.Values, definition models.Definitions, pdf *gofpdf.Fpdf, seal, logo []byte) {
	// write data to pdf
	for _, v := range definition.Texts {
		if v.Size == 0 {
			v.Size = 12
		}
		pdf.SetPage(v.PageNo)
		pdf.SetFont("Arial", "", float64(v.Size))
		pdf.SetXY(float64(v.X), float64(v.Y))
		for _, i := range value.Items {

			if i.FieldName == v.FieldName {

				pdf.Cell(40, 10, i.Value)
			}
		}
	}

	// write seal and logo to pdf
	for _, v := range definition.Images {
		if v.Name == "seal" {
			pdf.SetPage(v.PageNo)
			pdf.RegisterImageOptionsReader("seal", gofpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(seal))
			pdf.Image("seal", float64(v.X), float64(v.Y), float64(v.Width), float64(v.Height), false, "", 0, "")
		}
		if v.Name == "logo" {
			pdf.SetPage(v.PageNo)
			pdf.RegisterImageOptionsReader("logo", gofpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(logo))
			pdf.Image("logo", float64(v.X), float64(v.Y), float64(v.Width), float64(v.Height), false, "", 0, "")
		}
	}

	// write details to pdf
	for _, v := range value.Detail {
		pdf.SetPage(definition.Details.PageNo)
		pdf.SetFont("Arial", "", float64(definition.Details.Size))
		if len(v.Name) > 0 {
			pdf.SetXY(float64(definition.Details.Name.X), float64(definition.Details.Name.Y))
			pdf.Cell(40, 10, v.Name)
		}
		if v.Quantity > 0 {
			pdf.SetXY(float64(definition.Details.Quantity.X), float64(definition.Details.Quantity.Y))
			pdf.Cell(40, 10, strconv.Itoa(v.Quantity))
		}
		if v.Price > 0 {
			pdf.SetXY(float64(definition.Details.Price.X), float64(definition.Details.Price.Y))
			pdf.Cell(40, 10, strconv.Itoa(v.Price))
		}
		definition.Details.Name.Y += definition.Details.IncrementY
		definition.Details.Quantity.Y += definition.Details.IncrementY
		definition.Details.Price.Y += definition.Details.IncrementY
	}

}
