package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"

	//"github.com/aws/aws-sdk-go/service/s3"
)

func getSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"AKIA5UJLYIAI2AGYF3XH",
			"6T6F25VeSdwZQ0UOHfFp/EMu4xqoMaiO/27e2e+z",
			"", // a token will be created when the session it's used.
		),
	})
	if err != nil {
		panic(err)
	}

	return sess
}


func UploadFile(file multipart.File, fileHeader multipart.FileHeader) {
	session := getSession()
	uploader := s3manager.NewUploader(session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("api-rest-go-storage"),
		ACL: aws.String("public-read"),
		Key: aws.String(fileHeader.Filename),
		Body: file,
	})

	if err != nil {
		panic(err)
	}

}
