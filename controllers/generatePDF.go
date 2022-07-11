package controllers

import (
	"bytes"
	"context"
	"fmt"
	"grrow_pdf/models"
	"io/ioutil"
	"strconv"

	npdf "github.com/dslipak/pdf"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GeneratePDF(data models.RawData) []byte {

	//downloading files from aws ======================================
	temp, logoByte, sealByte := getFiles(data.Template)

	// pdf path =========================================================
	pdfPath := "temp.pdf"

	// saving the template file to storage
	ioutil.WriteFile(pdfPath, temp, 0644)

	// Create a new PDF document =========================================
	pdf := gofpdf.New("P", "mm", "A4", "")

	// get number of pages from template file ============================
	contents, err := npdf.Open(pdfPath)
	if err != nil {
		panic(err)
	}
	numberOfPages := contents.NumPage()

	// changing template received from api to readseeker==================
	// template := io.ReadSeeker(bytes.NewReader(temp))

	// Add a page to the document ========================================
	for i := 1; i <= numberOfPages; i++ {
		page := gofpdi.ImportPage(pdf, "temp.pdf", i, "/MediaBox")

		pdf.AddPage()

		// Draw imported template onto page
		gofpdi.UseImportedTemplate(pdf, page, 0, 0, 215, 0)
	}

	writeData(data, pdf, logoByte, sealByte)

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

	// delete the PDF file from aws storage
	deleteFile(data.Template)

	// delete the PDF file from local storage
	// err = os.Remove(pdfPath)
	// if err != nil {
	// 	panic(err)
	// }

	return pdfBytes
}

// downloading files from aws ===========================================
func getFiles(tempInt string) ([]byte, []byte, []byte) {

	svc := GetS3()

	// download template ================================================
	pdfTemplate := &s3.GetObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String(tempInt),
	}
	req1, out := svc.GetObject(context.TODO(), pdfTemplate)
	if out != nil {
		fmt.Println(out.Error())
	}
	// reading the template file received via API
	temp, _ := ioutil.ReadAll(req1.Body)

	// download logo =====================================================
	logo := &s3.GetObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String("logo.png"),
	}
	req2, out := svc.GetObject(context.TODO(), logo)
	if out != nil {
		fmt.Println(out.Error())
	}
	// reading the logo file received via API
	logoByte, _ := ioutil.ReadAll(req2.Body)

	// download seal =====================================================
	seal := &s3.GetObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String("seal.png"),
	}
	req3, out := svc.GetObject(context.TODO(), seal)
	if out != nil {
		fmt.Println(out.Error())
	}
	// reading the seal file received via API
	sealByte, _ := ioutil.ReadAll(req3.Body)

	defer req1.Body.Close()
	defer req2.Body.Close()
	defer req3.Body.Close()

	return temp, logoByte, sealByte

}

// delete file from s3 server ============================================
func deleteFile(key string) {
	svc := GetS3()
	input := &s3.DeleteObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String(key),
	}
	_, err := svc.DeleteObject(context.TODO(), input)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// write data on pdf
func writeData(pdfData models.RawData, pdf *gofpdf.Fpdf, logo, seal []byte) {

	final := pdf.PageNo()

	// print name
	pdf.SetPage(pdfData.Name.PageNo)
	pdf.SetFont("Helvetica", "", fontSize(pdfData.Name.Size))
	pdf.SetXY(float64(pdfData.Name.X), float64(pdfData.Name.Y))
	pdf.Cell(40, 10, pdfData.Name.Name)

	// print address
	pdf.SetPage(pdfData.Address.PageNo)
	pdf.SetFont("Helvetica", "", fontSize(pdfData.Address.Size))
	pdf.SetXY(float64(pdfData.Address.X), float64(pdfData.Address.Y))
	pdf.Cell(40, 10, pdfData.Address.Address)

	// print phone number
	pdf.SetPage(pdfData.PhoneNumber.PageNo)
	pdf.SetFont("Helvetica", "", fontSize(pdfData.PhoneNumber.Size))
	pdf.SetXY(float64(pdfData.PhoneNumber.X), float64(pdfData.PhoneNumber.Y))
	pdf.Cell(40, 10, strconv.Itoa(pdfData.PhoneNumber.PhoneNumber))

	// print zip code
	pdf.SetPage(pdfData.ZipAddress.PageNo)
	pdf.SetFont("Helvetica", "", fontSize(pdfData.ZipAddress.Size))
	pdf.SetXY(float64(pdfData.ZipAddress.X), float64(pdfData.ZipAddress.Y))
	pdf.Cell(40, 10, strconv.Itoa(pdfData.ZipAddress.ZipAddress))

	// print logo
	pdf.SetPage(pdfData.LogoData.PageNo)
	pdf.RegisterImageOptionsReader("logo", gofpdf.ImageOptions{ImageType: "png"}, bytes.NewReader(logo))
	pdf.Image("logo", float64(pdfData.LogoData.X), float64(pdfData.LogoData.Y), float64(pdfData.LogoData.Width), float64(pdfData.LogoData.Height), false, "", 0, "")

	// print seal
	pdf.SetPage(pdfData.SealData.PageNo)
	pdf.RegisterImageOptionsReader("seal", gofpdf.ImageOptions{ImageType: "png"}, bytes.NewReader(seal))
	pdf.Image("seal", float64(pdfData.SealData.X), float64(pdfData.SealData.Y), float64(pdfData.SealData.Width), float64(pdfData.SealData.Height), false, "", 0, "")

	// print products

	for _, product := range pdfData.Products {

		// print product name
		pdf.SetPage(product.ProductName.PageNo)
		pdf.SetFont("Helvetica", "", fontSize(product.ProductName.Size))
		pdf.SetXY(float64(product.ProductName.X), float64(product.ProductName.Y))
		pdf.Cell(40, 10, product.ProductName.Name)

		// print product quantity
		pdf.SetPage(product.ProductQuantity.PageNo)
		pdf.SetFont("Helvetica", "", fontSize(product.ProductQuantity.Size))
		pdf.SetXY(float64(product.ProductQuantity.X), float64(product.ProductQuantity.Y))
		pdf.Cell(40, 10, strconv.Itoa(product.ProductQuantity.Quantity))

		// print product price
		pdf.SetPage(product.ProductPrice.PageNo)
		pdf.SetFont("Helvetica", "", fontSize(product.ProductPrice.Size))
		pdf.SetXY(float64(product.ProductPrice.X), float64(product.ProductPrice.Y))
		pdf.Cell(40, 10, strconv.Itoa(product.ProductPrice.Price))
	}

	pdf.SetPage(final)

}

func fontSize(size int) float64 {

	fontSize := float64(12)

	if size != 0 {
		fontSize = float64(size)
	}

	return fontSize
}
