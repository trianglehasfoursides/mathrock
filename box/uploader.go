package box

import "github.com/aws/aws-sdk-go-v2/feature/s3/manager"

var Uploader = manager.NewUploader(Box)
