package api

import (
	"context"
	"encoding/json"
	"fmt"
	"grrow_pdf/controllers"
	"grrow_pdf/models"
	"net/http"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

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
// @Accept       application/json
// @Produce      json
// @Param        rawData body models.RawData true "rawData"
// @Success      200  {string}  models.RawData
// @Failure      400  {string}  error
// @Failure      404  {string}  error
// @Failure      500  {string}  error
// @Router       /addToTemplate [POST]
func AddToTemplate(res http.ResponseWriter, req *http.Request) {

	var newData models.RawData
	err := json.NewDecoder(req.Body).Decode(&newData)
	if err != nil {
		fmt.Println(err.Error())
	}

	if newData.Products == nil {
		products := req.FormValue("products")

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

	if newData.LogoData.Height == 0 {
		logoData := req.FormValue("logoData")
		logoDataStruct := models.LogoData{}
		err := json.Unmarshal([]byte(logoData), &logoDataStruct)
		if err != nil {
			fmt.Println(err.Error())
		}
		newData.LogoData = logoDataStruct
	}

	byteData := controllers.GeneratePDF(newData)
	println(byteData)

	// send response as json
	res.Write(byteData)

}

// upload template to S3 api
// @Summary      Upload template to S3
// @Description  Upload template to S3
// @Tags         UploadTemplate
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "file"
// @Success      200  {string}  message
// @Failure      400  {string}  error
// @Failure      404  {string}  error
// @Failure      500  {string}  error
// @Router       /uploadTemplate [POST]
func UploadTemplate(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	file, header, err := req.FormFile("file")
	// file, header, err := rawData.Request.FormFile("file")

	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error"))
		return
	}

	defer file.Close()

	filename := header.Filename

	svc := controllers.GetS3()

	uploader := manager.NewUploader(svc)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error"))
		return
	}

	fmt.Println(result)
	res.Write([]byte("success"))

}
