package storage

import (
	"context"
	"strings"
	"testing"
)

var mockFile = MockFile{
	bucket:      "123",
	objectName:  "/mock",
	file:        strings.NewReader("123"),
	fileSize:    3,
	contentType: "image/jpeg",
}

func TestNewDisk(t *testing.T) {
	_ = NewDisk("", "")
}

func Test_disk_Save(t *testing.T) {
	d := NewDisk("http://127.0.0.1:8888", "/tmp")
	type args struct {
		file FileObject
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1", args: args{file: mockFile}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := d.Save(nil, tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_disk_Share(t *testing.T) {
	d := NewDisk("http://127.0.0.1:8888/", "/tmp")
	type args struct {
		in0        context.Context
		in1        string
		objectName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1", args: args{
			in0:        nil,
			in1:        "",
			objectName: "hello",
		}, want: "http://127.0.0.1:8888/hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.Share(tt.args.in0, tt.args.in1, tt.args.objectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Share() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Share() got = %v, want %v", got, tt.want)
			}
		})
	}
}
