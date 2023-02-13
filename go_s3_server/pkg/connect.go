package pkg

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// @desp: Create new AWS session to interact with S3
func CreateSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return nil
	}
	return sess
}

// @desp: create new S3 client using session
func S3client() *s3.S3 {
	sess := CreateSession()
	svc := s3.New(sess)
	return svc
}

// @desp: Upload file to s3
func ClientUpload() {
	svc := S3client()
	// file := io.ReadSeeker()
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("go_s3_bucket"),
		Key:    aws.String("go_s3_object"),
		Body:   os.Stdin,
	})
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
}

func ClientDownload() {

	svc := S3client()
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("go_s3_bucket"),
		Key:    aws.String("go_s3_object"),
	})

	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File data:", string(data))
}
