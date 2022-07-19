package api

import (
	"context"
	"fmt"
	"grrow_pdf/controllers"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// upload images to S3 api
// @Summary      Upload images to S3
// @Description  Upload images to S3
// @Tags         Upload Images
// @Accept       multipart/form-data
// @Produce      json
// @Param        logo formData file true "logo"
// @Param        seal formData file true "seal"
// @Success      200  {string}  message
// @Failure      400  {string}  error
// @Failure      404  {string}  error
// @Failure      500  {string}  error
// @Router       /uploadImages [POST]
func UploadImages(res http.ResponseWriter, req *http.Request) {

	//cors
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//end cors

	req.ParseForm()

	// take out seal image and logo image from form
	//seal
	sealImage, header, err := req.FormFile("seal")
	if err != nil {
		fmt.Println(err.Error())
	}
	sealImageFileName := strconv.Itoa(int(time.Now().Unix())) + "_" + header.Filename

	// logo
	logoImage, header, err := req.FormFile("logo")
	if err != nil {
		fmt.Println(err.Error())
	}
	logoImageFileName := strconv.Itoa(int(time.Now().Unix())) + "_" + header.Filename

	//aws ===========================================================

	svc := controllers.GetS3()

	//upload seal image
	uploader := manager.NewUploader(svc)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String(sealImageFileName),
		Body:   sealImage,
	})
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error upload seal image"))
		return
	}
	fmt.Println(result)
	fmt.Println("seal image uploaded")

	//upload logo image
	uploader = manager.NewUploader(svc)
	result, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("grrow.pdf.generator"),
		Key:    aws.String(logoImageFileName),
		Body:   logoImage,
	})
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error upload logo image"))
		return
	}
	fmt.Println(result)
	fmt.Println("logo image uploaded")

	returnString := "{\"sealImage\":\"" + sealImageFileName + "\",\"logoImage\":\"" + logoImageFileName + "\"}"

	res.Write([]byte(returnString))
}
