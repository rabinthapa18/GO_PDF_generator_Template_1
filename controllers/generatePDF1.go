package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	models "grrow_pdf/models"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// pdf generation code
func GeneratePDF1(pdfData models.PdfData) {
	tax := 0
	withholdingTax := 0
	finalAmount := 0

	// get total amount
	total := setTotalAmount(pdfData.Products)
	tax = int(float64(total) * 1.1)
	withholdingTax = int(float64(total) * 0.121)
	finalAmount = tax - withholdingTax

	pdf := pdf.NewMaroto(consts.Portrait, consts.A4)
	pdf.SetPageMargins(20, 15, 20)
	// pdf.SetBorder(true)

	// changing image to base 64
	data, _ := ioutil.ReadAll(pdfData.Logo)

	contentType := http.DetectContentType(data)
	base64Image := base64.StdEncoding.EncodeToString(data)
	base64Image += string(contentType)

	//adding fonts
	pdf.AddUTF8Font("ZenKakuGothicAntique", consts.Normal, "/fonts/ZenKakuGothicAntique-Regular.ttf")
	pdf.AddUTF8Font("ZenKakuGothicAntique", consts.Bold, "/fonts/ZenKakuGothicAntique-Bold.ttf")

	//building components of pdf
	buildHeader(pdf)
	buildNameAddressPhone(pdf, pdfData.Name, pdfData.Address, (pdfData.PhoneNumber), pdfData.ZipAddress, base64Image)
	buildSubHeader(pdf, pdfData.Name)
	whitespace(pdf, 20)
	buildTotalAmountPaymentDate(finalAmount, pdf)
	whitespace(pdf, 20)
	buildProductsTable(pdf, pdfData.Products)
	whitespace(pdf, 3)
	buildBillingAmount(pdf, total, tax, withholdingTax, finalAmount)
	whitespace(pdf, 10)
	buildTransferAddress(pdf)

	err := pdf.OutputFileAndClose("pdfs/template1trial.pdf")
	if err != nil {
		fmt.Println(err)
	}

}

// header ( invoice heading and date time)
func buildHeader(pdf pdf.Maroto) {
	pdf.Row(10, func() {
		pdf.Col(7, func() {
			// invoice header
			pdf.Text("請求書", props.Text{
				Family: "ZenKakuGothicAntique",
				Size:   20,
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
		pdf.Col(5, func() {
			//date
			pdf.Row(10, func() {
				pdf.Col(12, func() {
					pdf.Text("Date", props.Text{
						Top:   3,
						Style: consts.Bold,
						Align: consts.Right,
						Color: blackColor(),
					})
				})
			})
			pdf.Row(10, func() {
				pdf.Col(12, func() {

					t := time.Now().Format("2006-01-02 15:04:05")
					pdf.Text(t, props.Text{

						Style: consts.Bold,
						Align: consts.Right,
						Color: blackColor(),
					})
				})
			})
		})
	})

}

// name address and phone number fields and logo
func buildNameAddressPhone(pdf pdf.Maroto, name, address string, phone, zip int, logo string) {

	// name
	pdf.Row(8, func() {
		pdf.ColSpace(9)
		pdf.Text(name, props.Text{
			Family: "ZenKakuGothicAntique",
			Style:  consts.Bold,
			Size:   12,
			Color:  blackColor(),
		})
	})

	// zip code
	pdf.Row(5, func() {
		pdf.ColSpace(9)
		pdf.Text("〒"+strconv.Itoa(zip), props.Text{
			Family: "ZenKakuGothicAntique",
			Style:  consts.Bold,
			Color:  blackColor(),
		})
	})

	// address
	pdf.Row(14, func() {
		pdf.ColSpace(9)
		pdf.Col(3, func() {
			pdf.Text(address, props.Text{
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})

	})

	// phone number
	pdf.Row(3, func() {
		pdf.ColSpace(9)
		pdf.Text(strconv.Itoa(phone), props.Text{
			Family: "ZenKakuGothicAntique",
			Style:  consts.Bold,
			Color:  blackColor(),
		})
	})

	// logo
	pdf.Row(10, func() {
		pdf.ColSpace(11)
		pdf.Base64Image(logo, consts.Png)
	})
}

// whitespace
func whitespace(pdf pdf.Maroto, rows float64) {
	pdf.SetBorder(false)
	pdf.Row(rows, func() {
	})
}

// sub header (name and text)
func buildSubHeader(pdf pdf.Maroto, name string) {
	pdf.Col(6, func() {
		pdf.Row(5, func() {
			pdf.Text(name, props.Text{
				Family: "ZenKakuGothicAntique",
				Align:  consts.Center,
				Style:  consts.Bold,
				Size:   12,
				Color:  blackColor(),
			})
		})
		pdf.Row(8, func() {
			pdf.Text("__________________________________________________", props.Text{
				Style: consts.Bold,
			})
		})
		pdf.Row(5, func() {
			pdf.Text("下記の通りご請求申し上げます。", props.Text{
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Size:   12,
				Color:  blackColor(),
			})
		})
	})

}

// total amount and payment date
func buildTotalAmountPaymentDate(finalAmount int, pdf pdf.Maroto) {

	// text area
	pdf.Row(5, func() {
		pdf.Col(3, func() {
			pdf.Text("ご請求金（税込）", props.Text{
				Top:    3,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Size:   14,
				Color:  blackColor(),
			})
		})
		pdf.Col(4, func() {
			pdf.Text(" ¥ "+strconv.Itoa(finalAmount), props.Text{
				Top:    0,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Size:   22,
				Color:  blackColor(),
			})

		})
		pdf.Col(2, func() {
			pdf.Text("お支払期日", props.Text{
				Top:    3,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Size:   14,
				Color:  blackColor(),
			})
		})
		pdf.Col(3, func() {
			pdf.Text("2022-06-18", props.Text{
				Family: "ZenKakuGothicAntique",
				Align:  consts.Right,
				Style:  consts.Bold,
				Size:   18,
				Color:  blackColor(),
			})
		})
	})

	// unerline area
	pdf.Row(1, func() {
		pdf.ColSpace(3)
		pdf.Col(4, func() {
			pdf.Text("______________________", props.Text{
				Style: consts.Bold,
			})
		})
		pdf.ColSpace(2)
		pdf.Col(3, func() {
			pdf.Text("____________________", props.Text{
				Style: consts.Bold,
				Align: consts.Right,
			})
		})
	})

}

// products table
func buildProductsTable(pdf pdf.Maroto, products []models.ProductData) {
	pdf.SetBorder(true)

	// product list
	contents := [][]string{}
	for i := 0; i < len(products); i++ {
		contents = append(
			contents,
			[]string{
				" " + products[i].ProductName,
				"  " + strconv.Itoa(products[i].Quantity),
				" ¥ " + strconv.Itoa(products[i].Price),
				"  ¥ " + strconv.Itoa(products[i].Price*products[i].Quantity)})
	}

	tableHeadings := []string{" 品名", "", "額", "金額（税抜)"}

	// table header
	pdf.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      14,
			Family:    "ZenKakuGothicAntique",
			GridSizes: []uint{7, 1, 2, 2},
		},
		ContentProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{7, 1, 2, 2},
			Family:    "ZenKakuGothicAntique",
		},
		HeaderContentSpace:     1,
		VerticalContentPadding: 3,
	})

}

// billing amount
func buildBillingAmount(pdf pdf.Maroto, total, tax, withholdingTax, finalAmount int) {
	pdf.SetBorder(true)

	// total amount
	pdf.Row(8, func() {
		pdf.Col(10, func() {
			pdf.Text("小 計 ", props.Text{
				Top:    2,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
		pdf.Col(2, func() {
			pdf.Text("¥ "+strconv.Itoa(total), props.Text{
				Top:    2,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
	})

	// tax amount
	pdf.Row(8, func() {
		pdf.Col(10, func() {
			pdf.Text("税込小計（10％）", props.Text{
				Top:    2,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
		pdf.Col(2, func() {
			pdf.Text("¥ "+fmt.Sprintf("%d", int(tax)), props.Text{
				Top:    2,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
	})

	// tax withholding
	pdf.Row(8, func() {
		pdf.Col(10, func() {
			pdf.Text("源泉徴収（10.21％） ", props.Text{
				Top: 2,

				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
		pdf.Col(2, func() {
			pdf.Text("¥ "+fmt.Sprintf("%d", int(withholdingTax)), props.Text{
				Top:    2,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
	})

	// final amount
	pdf.Row(8, func() {
		pdf.Col(10, func() {
			pdf.Text("請求金額合計 ", props.Text{
				Top:    2,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
		pdf.Col(2, func() {
			pdf.Text("¥ "+fmt.Sprintf("%d", int(finalAmount)), props.Text{
				Top:    2,
				Family: "ZenKakuGothicAntique",
				Style:  consts.Bold,
				Color:  blackColor(),
			})
		})
	})

}

// transfer address
func buildTransferAddress(pdf pdf.Maroto) {

	pdf.SetBorder(true)

	pdf.Row(40, func() {
		pdf.Col(12, func() {
			pdf.Text("■振込先", props.Text{
				Top:    1,
				Family: "ZenKakuGothicAntique",
				Color:  blackColor(),
			})
		})
	})
}

// color
func blackColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
}

// set total amount
func setTotalAmount(products []models.ProductData) int {
	total := 0
	for i := 0; i < len(products); i++ {
		total += products[i].Price * products[i].Quantity
	}
	return total
}
