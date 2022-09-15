package storage

import (
	"context"
	"io"
)

type FileObject interface {
	GetBucket() string
	GetObjectName() string
	GetFile() io.Reader
	GetFileSize() int64
	GetContentType() string
}

type Storage interface {
	Save(ctx context.Context, file FileObject) error
	Share(ctx context.Context, bucket string, objectName string) (string, error)
}
