package api

import (
	"encoding/json"
	"fmt"
	"grrow_pdf/controllers"
	"grrow_pdf/models"
	"reflect"

	"github.com/gin-gonic/gin"
)

// func GetData(pdfData *gin.Context) {
// 	pdfData.IndentedJSON(http.StatusOK, data)
// }

// generate PDF api
// @Summary      Create new pdf
// @Description  Create new pdf
// @Tags         GeneratePDF
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
func AddData(pdfData *gin.Context) {

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
