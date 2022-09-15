package utils

import (
	"net"
	"testing"
)

func TestIP(t *testing.T) {
	type args struct {
		neighbor net.IP
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1", args: args{neighbor: nil}},
		{name: "1", args: args{neighbor: nil}},
		{name: "2", args: args{neighbor: net.IPv4(1, 2, 3, 4)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := IP(tt.args.neighbor)
			if (err != nil) != tt.wantErr {
				t.Errorf("IP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
