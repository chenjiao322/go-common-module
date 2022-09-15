package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"time"
)

type Minio struct {
	client *minio.Client
}

func NewMinio(endpoint string, option *minio.Options) (*Minio, error) {
	c, err := minio.New(endpoint, option)
	return &Minio{client: c}, err
}

func (m Minio) Save(ctx context.Context, file FileObject) error {
	_, err := m.client.PutObject(
		ctx,
		file.GetBucket(),
		file.GetObjectName(),
		file.GetFile(),
		file.GetFileSize(),
		minio.PutObjectOptions{ContentType: file.GetContentType()})
	return err
}

func (m Minio) Share(ctx context.Context, bucket string, objectName string) (string, error) {
	u, err := m.client.PresignedGetObject(ctx, bucket, objectName, time.Hour*24, nil)
	return u.String(), err
}
