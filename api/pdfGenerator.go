package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grrow_pdf/controllers"
	"grrow_pdf/models"
	"io/ioutil"
	"net/http"
	"os/exec"
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
	//cors
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//end cors

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

	//cors
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//end cors

	req.ParseForm()
	file, header, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error"))
		return
	}

	//change file to byte
	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error"))
		return
	}

	// save file to local storage
	ioutil.WriteFile("temp.pdf", byteData, 0644)
	ioutil.WriteFile("temp2.pdf", byteData, 0644)

	// cmd := exec.Command("gswin64c", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4", "-dPDFSETTINGS=/screen", "-dNOPAUSE", "-dQUIET", "-dBATCH", "-sOutputFile=temp2.pdf", "temp.pdf")
	// cmd.Run()

	cmd := exec.Command("gs", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4", "-dPDFSETTINGS=/screen", "-dNOPAUSE", "-dQUIET", "-dBATCH", "-sOutputFile=temp2.pdf", "temp.pdf")
	cmd.Run()

	pdfByte, err := ioutil.ReadFile("temp2.pdf")
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error"))
		return
	}

	filename := header.Filename

	svc := controllers.GetS3()

	uploader := manager.NewUploader(svc)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(pdfByte),
	})

	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error"))
		return
	}

	fmt.Println(result)

	cmd = exec.Command("rm", "temp.pdf")
	cmd.Run()
	cmd = exec.Command("rm", "temp2.pdf")
	cmd.Run()

	res.Write([]byte("success"))

}
