package s3

import (
	"DATN/configs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type S3Repo struct {
	cfg      *configs.Server
	uploader *s3manager.Uploader
}

func NewS3Repo(cfg *configs.Server) (IS3Repo, error) {
	conf := &aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.Id, cfg.Secret, ""),
		Region:      aws.String(cfg.Region),
	}
	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, err
	}
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024
		u.Concurrency = 5
		u.LeavePartsOnError = false
	})
	return &S3Repo{
		cfg:      cfg,
		uploader: uploader,
	}, nil
}

func (s *S3Repo) PutObject(path string, body *os.File) error {
	input := &s3manager.UploadInput{
		Body:   body,
		Bucket: aws.String(s.cfg.Bucket),
		Key:    aws.String(path),
	}
	_, err := s.uploader.Upload(input)
	if err != nil {
		return err
	}
	return nil
}
