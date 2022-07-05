package controllers

import (
	"bytes"
	"context"
	"fmt"
	"grrow_pdf/models"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	npdf "github.com/dslipak/pdf"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GeneratePDF(data models.RawData) []byte {

	// pdf path =========================================================
	pdfPath := "pdfs/PDF_" + time.Now().Format("2006-01-02_15-04-05") + ".pdf"

	//downloading template from aws ======================================
	svc := GetS3()
	input := &s3.GetObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String("pdf-template-" + strconv.Itoa(data.Template) + ".pdf"),
	}
	req, out := svc.GetObject(context.TODO(), input)
	if out != nil {
		fmt.Println(out.Error())
	}
	defer req.Body.Close()

	// reading the template file received via API ========================
	temp, _ := ioutil.ReadAll(req.Body)

	// saving the template file to storage
	ioutil.WriteFile(pdfPath, temp, 0644)

	// Create a new PDF document =========================================
	pdf := gofpdf.New("P", "mm", "A4", "")

	// get number of pages from template file ============================
	c, e := npdf.NewReader(bytes.NewReader(temp), 0644)
	if e != nil {
		fmt.Println(e)
	}
	pages := c.NumPage()
	fmt.Println(pages)
	contents, err := npdf.Open(pdfPath)
	if err != nil {
		panic(err)
	}
	numberOfPages := contents.NumPage()

	// changing template received from api to readseeker==================
	if err != nil {
		panic(err)
	}
	template := io.ReadSeeker(bytes.NewReader(temp))

	// Add a page to the document ========================================
	for i := 1; i <= numberOfPages; i++ {
		page := gofpdi.ImportPageFromStream(pdf, &template, 1, "/MediaBox")

		pdf.AddPage()

		// Draw imported template onto page
		gofpdi.UseImportedTemplate(pdf, page, 0, 0, 215, 0)
	}

	writeData(data, pdf)

	// Output the document ==============================================
	// buff := pdf.Output()

	// Output the document to a file =====================================
	err = pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		panic(err)
	}

	// change generated PDF to bytes
	pdfBytes, err := ioutil.ReadFile(pdfPath)
	if err != nil {
		panic(err)
	}

	// delete the PDF file from local storage
	err = os.Remove(pdfPath)
	if err != nil {
		panic(err)
	}

	return pdfBytes
}

// write data on pdf
func writeData(pdfData models.RawData, pdf *gofpdf.Fpdf) {
	// print name
	pdf.SetFont("Helvetica", "", fontSize(pdfData.Name.Size))
	pdf.SetXY(float64(pdfData.Name.X), float64(pdfData.Name.Y))
	pdf.Cell(40, 10, pdfData.Name.Name)

	// print address
	pdf.SetFont("Helvetica", "", fontSize(pdfData.Address.Size))
	pdf.SetXY(float64(pdfData.Address.X), float64(pdfData.Address.Y))
	pdf.Cell(40, 10, pdfData.Address.Address)

	// print phone number
	pdf.SetFont("Helvetica", "", fontSize(pdfData.PhoneNumber.Size))
	pdf.SetXY(float64(pdfData.PhoneNumber.X), float64(pdfData.PhoneNumber.Y))
	pdf.Cell(40, 10, strconv.Itoa(pdfData.PhoneNumber.PhoneNumber))

	// print zip code
	pdf.SetFont("Helvetica", "", fontSize(pdfData.ZipAddress.Size))
	pdf.SetXY(float64(pdfData.ZipAddress.X), float64(pdfData.ZipAddress.Y))
	pdf.Cell(40, 10, strconv.Itoa(pdfData.ZipAddress.ZipAddress))

	// print logo
	pdf.RegisterImageOptionsReader("logo", gofpdf.ImageOptions{ImageType: "png"}, ioutil.NopCloser(pdfData.Logo))
	pdf.Image("logo", float64(pdfData.LogoData.X), float64(pdfData.LogoData.Y), float64(pdfData.LogoData.Width), float64(pdfData.LogoData.Height), false, "", 0, "")

	// print products

	for _, product := range pdfData.Products {

		// print product name
		pdf.SetFont("Helvetica", "", fontSize(product.ProductName.Size))
		pdf.SetXY(float64(product.ProductName.X), float64(product.ProductName.Y))
		pdf.Cell(40, 10, product.ProductName.Name)

		// print product quantity
		pdf.SetFont("Helvetica", "", fontSize(product.ProductQuantity.Size))
		pdf.SetXY(float64(product.ProductQuantity.X), float64(product.ProductQuantity.Y))
		pdf.Cell(40, 10, strconv.Itoa(product.ProductQuantity.Quantity))

		// print product price
		pdf.SetFont("Helvetica", "", fontSize(product.ProductPrice.Size))
		pdf.SetXY(float64(product.ProductPrice.X), float64(product.ProductPrice.Y))
		pdf.Cell(40, 10, strconv.Itoa(product.ProductPrice.Price))
	}

}

func fontSize(size int) float64 {

	fontSize := float64(12)

	if size != 0 {
		fontSize = float64(size)
	}

	return fontSize
}
