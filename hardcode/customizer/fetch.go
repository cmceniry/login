package customizer

import (
	"strings"

	// tag::imports[]
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	// end::imports[]
)

// tag::fetch[]
func fetch(ak, sk, bucket, path string) error {
	// end::fetch[]

	// tag::session[]
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(
			ak, sk, "",
		),
	})
	// end::session[]
	if err != nil {
		return err
	}

	// tag::writeatbuf[]
	writeAtBuf := aws.NewWriteAtBuffer([]byte{})
	// end::writeatbuf[]
	// tag::download[]
	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(
		writeAtBuf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(path),
		},
	)
	// end::download[]
	if err != nil {
		return err
	}

	// tag::config[]
	globalConfig = strings.TrimSpace(
		string(
			writeAtBuf.Bytes(),
		),
	)
	// end::config[]

	return nil
}
