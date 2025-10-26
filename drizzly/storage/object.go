package storage

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type object struct {
	client   *s3.Client
	uploader *manager.Uploader
}

func (o *object) Setup() (err error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID {
			return aws.Endpoint{
				URL:           os.Getenv(""),
				SigningRegion: "us-east-1", // Region bisa disesuaikan
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv(""), os.Getenv(""), "")),
	)

	if err != nil {
		return
	}

	o.client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true // Penting untuk MinIO
	})

	o.uploader = manager.NewUploader(o.client)

	return
}

func (o *object) Upload(file *os.File) (err error) {
	_, err = o.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(""),
		Body:   file,
	})

	return
}

func (o *object) Get(name string) (err error) {
	_, err = o.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(""),
	})

	return
}

func (o *object) Delete(name string) (err error) {
	_, err = o.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(""),
	})

	return
}
