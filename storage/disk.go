package storage

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type disk struct {
	Dir  string
	Host string
}

func NewDisk(host, dir string) *disk {
	return &disk{Dir: dir, Host: host}
}

func (d disk) Save(_ context.Context, file FileObject) error {
	absPath := filepath.Join(d.Dir, file.GetObjectName())
	err := os.MkdirAll(filepath.Dir(absPath), 0664)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(absPath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0664)
	defer func() { _ = f.Close() }()
	n, err := io.Copy(f, file.GetFile())
	if n != file.GetFileSize() {
		return errors.New("save to disk break")
	}
	return err
}

func (d disk) Share(_ context.Context, _ string, objectName string) (string, error) {
	return d.Host + objectName, nil
}
