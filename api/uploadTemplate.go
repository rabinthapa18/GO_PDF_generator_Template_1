package api

import (
	"bytes"
	"context"
	"fmt"
	"grrow_pdf/controllers"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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

	cmd := exec.Command("gswin64c", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4", "-dPDFSETTINGS=/screen", "-dNOPAUSE", "-dQUIET", "-dBATCH", "-sOutputFile=temp2.pdf", "temp.pdf")
	// cmd.Run()

	// cmd := exec.Command("gs", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4", "-dPDFSETTINGS=/screen", "-dNOPAUSE", "-dQUIET", "-dBATCH", "-sOutputFile=temp2.pdf", "temp.pdf")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("error"))
		return
	}

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
