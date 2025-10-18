package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/trianglehasfoursides/mathrock/box"
)

var bucket *s3.CreateBucketOutput

func SetupBucket() (err error) {
	bucket, err = box.Box.CreateBucket(
		context.Background(),
		&s3.CreateBucketInput{
			Bucket: aws.String("storage"),
		},
	)

	return
}
