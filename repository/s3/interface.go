package s3

import "os"

type IS3Repo interface {
	PutObject(path string, body *os.File) error
}
