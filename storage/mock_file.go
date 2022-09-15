package storage

import (
	"io"
)

// MockFile 测试用
type MockFile struct {
	bucket      string
	objectName  string
	file        io.Reader
	fileSize    int64
	contentType string
}

func (m MockFile) GetBucket() string {
	return m.bucket
}

func (m MockFile) GetObjectName() string {
	return m.objectName
}

func (m MockFile) GetFile() io.Reader {
	return m.file
}

func (m MockFile) GetFileSize() int64 {
	return m.fileSize
}

func (m MockFile) GetContentType() string {
	return m.contentType
}
