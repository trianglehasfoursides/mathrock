package box

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var Box *s3.Client

func Setup() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			credentials.
				NewStaticCredentialsProvider(
					os.Getenv(""),
					os.Getenv(""),
					"",
				),
		),
	)

	if err != nil {
		return err
	}

	Box = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.EndpointResolver = s3.EndpointResolverFromURL("http://localhost:9000")
		o.UsePathStyle = true // penting! untuk MinIO
	})

	return nil
}
