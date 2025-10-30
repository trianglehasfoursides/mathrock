package storage

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.etcd.io/bbolt"
)

type object struct {
	client   *s3.Client
	uploader *manager.Uploader
}

func (o *object) Setup() (err error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID {
			return aws.Endpoint{
				URL: os.Getenv(""),
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

// TODO:
func (o *object) Upload(fhead *multipart.FileHeader) (hashstr string, err error) {
	file, err := fhead.Open()
	if err != nil {
		return
	}

	hash := sha256.New()
	hashsum := hash.Sum(nil)
	hashstr = hex.EncodeToString(hashsum)

	if err = metatable.View(func(tx *bbolt.Tx) (err error) {
		if tx.Bucket([]byte("hashes")).Get([]byte(hashstr)) == nil {
			return errors.New("")
		}

		return
	}); err != nil {
		return
	}

	if _, err = o.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(""),
		Body:   file,
	}); err != nil {
		return
	}

	metatable.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("hashes")).Put([]byte(hashstr), []byte(fhead.Filename))
	})

	return
}

// TODO
func (o *object) Get(hash string) (file io.ReadCloser, err error) {
	if err = metatable.View(func(tx *bbolt.Tx) (err error) {
		if tx.Bucket([]byte("")).Get([]byte(hash)) != nil {
			return
		}

		return
	}); err != nil {
		return
	}

	obj, err := o.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(hash),
	})

	if err != nil {
		return
	}

	file = obj.Body
	return
}

// TODO
func (o *object) Delete(name string) (err error) {
	_, err = o.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(""),
	})

	return
}
