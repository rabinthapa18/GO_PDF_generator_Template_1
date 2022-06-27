package api

import (
	"encoding/json"
	"fmt"
	"grrow_pdf/controllers"
	"grrow_pdf/models"
	"reflect"

	"github.com/gin-gonic/gin"
)

// generate PDF with template api
// @Summary      Create new pdf from scratch
// @Description  Create new pdf from scratch
// @Tags         GeneratePDF1
// @Accept       multipart/form-data
// @Produce      json
// @Param        pdfData formData models.PdfData true "pdfData"
// @Param		 products formData []string true "products"
// @Param 	  	 logo formData file true "logo"
// @Success      200  {string}  models.PdfData
// @Failure      400  {string}  error
// @Failure      404  {string}  error
// @Failure      500  {string}  error
// @Router       /addData [POST]
func GenerateTemp1(pdfData *gin.Context) {

	var newData models.PdfData
	pdfData.ShouldBind(&newData)

	if newData.Products == nil {
		products := pdfData.PostForm("products")

		isString := reflect.TypeOf(products) == reflect.TypeOf(" ")
		if isString {
			products = "[" + products + "]"
			productStruct := []models.ProductData{}
			err := json.Unmarshal([]byte(products), &productStruct)
			if err != nil {
				fmt.Println(err.Error())
			}
			newData.Products = productStruct
		}
	}

	file, _, _ := pdfData.Request.FormFile("logo")

	newData.Logo = file

	controllers.GeneratePDF1(newData)

}

// generate PDF with position on template api
// @Summary      Create new pdf with positions
// @Description  Create new pdf with positions
// @Tags         GeneratePDF
// @Accept       multipart/form-data
// @Produce      json
// @Param        rawData formData models.RawData true "rawData"
// @Param        name formData object true "name"
// @Param		 phoneNumber formData object true "phoneNumber"
// @Param 		 zipAddress formData object true "zipAddress"
// @Param 		 address formData object true "address"
// @Param 		 logoData formData object true "logoData"
// @Param		 products formData []string true "products"
// @Param 	  	 logo formData file true "logo"
// @Param 	  	 template formData file true "template"
// @Success      200  {string}  models.RawData
// @Failure      400  {string}  error
// @Failure      404  {string}  error
// @Failure      500  {string}  error
// @Router       /addToTemplate [POST]
func AddToTemplate(rawData *gin.Context) {

	var newData models.RawData
	rawData.ShouldBind(&newData)

	if newData.Products == nil {
		products := rawData.PostForm("products")

		isString := reflect.TypeOf(products) == reflect.TypeOf(" ")
		if isString {

			products = "[" + products + "]"
			productStruct := []models.Products{}

			err := json.Unmarshal([]byte(products), &productStruct)
			if err != nil {
				fmt.Println(err.Error())
			}
			newData.Products = productStruct
		}

	}

	logo, _, _ := rawData.Request.FormFile("logo")
	newData.Logo = logo

	template, _, _ := rawData.Request.FormFile("template")
	newData.Template = template

	if newData.LogoData.Height == 0 {
		logoData := rawData.PostForm("logoData")
		logoDataStruct := models.LogoData{}
		err := json.Unmarshal([]byte(logoData), &logoDataStruct)
		if err != nil {
			fmt.Println(err.Error())
		}
		newData.LogoData = logoDataStruct
	}

	controllers.GeneratePDF(newData)

}