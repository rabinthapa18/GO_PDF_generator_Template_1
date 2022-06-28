package controllers

import (
	"bytes"
	"grrow_pdf/models"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	npdf "github.com/dslipak/pdf"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func GeneratePDF(data models.RawData) []byte {
	// Create a new PDF document =========================================
	pdf := gofpdf.New("P", "mm", "A4", "")

	// reading the template file received via API
	temp, _ := ioutil.ReadAll(data.Template)

	templateName := strconv.Itoa(int(time.Now().UnixNano()))

	// saving the template file to storage
	ioutil.WriteFile("templates/"+templateName+".pdf", temp, 0644)

	// get number of pages from template file ============================
	contents, err := npdf.Open("templates/pdf-template-1.pdf")
	if err != nil {
		panic(err)
	}
	numberOfPages := contents.NumPage()

	// changing template received from api to readseeker
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

	// Output the document to a file =====================================
	fileName := strconv.Itoa(int(time.Now().UnixNano()))
	err = pdf.OutputFileAndClose("pdfs/" + fileName + ".pdf")
	if err != nil {
		panic(err)
	}

	// delete the template file from storage
	err = os.Remove("templates/" + templateName + ".pdf")
	if err != nil {
		panic(err)
	}

	// change filename.pdf to bytes
	pdfBytes, err := ioutil.ReadFile("pdfs/" + fileName + ".pdf")
	if err != nil {
		panic(err)
	}
	return pdfBytes
}

// write data on pdf
func writeData(pdfData models.RawData, pdf *gofpdf.Fpdf) {
	// print name
	pdf.SetFont("Helvetica", "", 20)
	pdf.SetXY(float64(pdfData.Name.X), float64(pdfData.Name.Y))
	pdf.Cell(40, 10, pdfData.Name.Name)

	// print address
	pdf.SetFont("Helvetica", "", 10)
	pdf.SetXY(float64(pdfData.Address.X), float64(pdfData.Address.Y))
	pdf.Cell(40, 10, pdfData.Address.Address)

	// print phone number
	pdf.SetFont("Helvetica", "", 10)
	pdf.SetXY(float64(pdfData.PhoneNumber.X), float64(pdfData.PhoneNumber.Y))
	pdf.Cell(40, 10, strconv.Itoa(pdfData.PhoneNumber.PhoneNumber))

	// print zip code
	pdf.SetFont("Helvetica", "", 10)
	pdf.SetXY(float64(pdfData.ZipAddress.X), float64(pdfData.ZipAddress.Y))
	pdf.Cell(40, 10, strconv.Itoa(pdfData.ZipAddress.ZipAddress))

	// print logo
	pdf.RegisterImageOptionsReader("logo", gofpdf.ImageOptions{ImageType: "png"}, ioutil.NopCloser(pdfData.Logo))
	pdf.Image("logo", float64(pdfData.LogoData.X), float64(pdfData.LogoData.Y), float64(pdfData.LogoData.Width), float64(pdfData.LogoData.Height), false, "", 0, "")

	// print products
	pdf.SetFont("Helvetica", "", 8)

	for _, product := range pdfData.Products {

		// print product name
		pdf.SetXY(float64(product.ProductName.X), float64(product.ProductName.Y))
		pdf.Cell(40, 10, product.ProductName.Name)

		// print product quantity
		pdf.SetXY(float64(product.ProductQuantity.X), float64(product.ProductQuantity.Y))
		pdf.Cell(40, 10, strconv.Itoa(product.ProductQuantity.Quantity))

		// print product price
		pdf.SetXY(float64(product.ProductPrice.X), float64(product.ProductPrice.Y))
		pdf.Cell(40, 10, strconv.Itoa(product.ProductPrice.Price))
	}

}
