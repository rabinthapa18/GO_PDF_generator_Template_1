package controllers

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	awsS3Client *s3.Client
	prefix      string
	delimeter   string
)

const (
	BUCKET_NAME = "grrow.pdf.generator"
	REGION      = "ap-northeast-1"
)

type S3ListBucketsAPI interface {
	ListBuckets(ctx context.Context,
		params *s3.ListBucketsInput,
		optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func GetAllBuckets(c context.Context, api S3ListBucketsAPI, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return api.ListBuckets(c, input)
}

func GetS3() (awsS3Client *s3.Client) {
	// env.Config()

	creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(REGION))
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	awsS3Client = s3.NewFromConfig(cfg)

	return awsS3Client

	// paginator := s3.NewListObjectsV2Paginator(awsS3Client, &s3.ListObjectsV2Input{
	// 	Bucket:    aws.String(BUCKET_NAME),
	// 	Prefix:    aws.String(prefix),
	// 	Delimiter: aws.String(delimeter),
	// })

	// for paginator.HasMorePages() {
	// 	page, _ := paginator.NextPage(context.TODO())
	// 	for _, obj := range page.Contents {
	// 		// Do whatever you need with each object "obj"
	// 		fmt.Println("Content")
	// 		fmt.Println(obj)
	// 	}
	// }
}
