package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"testing"
)

func TestMinio_Save(t *testing.T) {
	type fields struct {
		client *minio.Client
	}
	type args struct {
		ctx  context.Context
		file FileObject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Minio{
				client: tt.fields.client,
			}
			if err := m.Save(tt.args.ctx, tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinio_Share(t *testing.T) {
	type fields struct {
		client *minio.Client
	}
	type args struct {
		ctx        context.Context
		bucket     string
		objectName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Minio{
				client: tt.fields.client,
			}
			got, err := m.Share(tt.args.ctx, tt.args.bucket, tt.args.objectName)
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

func TestNewMinio(t *testing.T) {
	_, _ = NewMinio("", nil)
}
