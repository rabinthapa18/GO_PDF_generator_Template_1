package aws_s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteObject(filename string, s3session *s3.S3, BUCKET_NAME string) (resp *s3.DeleteObjectOutput) {
	fmt.Println("Deleting: ", filename)
	resp, err := s3session.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
	})

	if err != nil {
		panic(err)
	}

	return resp
}
