package controllers

import (
	"grrow_pdf/models"
	"io/ioutil"
	"strconv"

	npdf "github.com/dslipak/pdf"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func GeneratePDF(data models.RawData) {
	// Create a new PDF document =========================================
	pdf := gofpdf.New("P", "mm", "A4", "")

	// get number of pages from template file ============================
	contents, err := npdf.Open("templates/pdf-template-1.pdf")
	if err != nil {
		panic(err)
	}

	numberOfPages := contents.NumPage()

	// Add a page to the document ========================================
	for i := 1; i <= numberOfPages; i++ {
		// Import example-pdf.pdf with gofpdi free pdf document importer
		page := gofpdi.ImportPage(pdf, "templates/pdf-template-1.pdf", i, "/MediaBox")

		pdf.AddPage()

		// Draw imported template onto page
		gofpdi.UseImportedTemplate(pdf, page, 0, 0, 215, 0)
	}

	writeData(data, pdf)

	err = pdf.OutputFileAndClose("pdfs/t1.pdf")
	if err != nil {
		panic(err)
	}
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
