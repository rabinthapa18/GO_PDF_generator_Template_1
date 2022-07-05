package aws_s3

import "github.com/aws/aws-sdk-go/service/s3"

func ListBuckets(s3session *s3.S3) (resp *s3.ListBucketsOutput) {
	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	return resp
}
