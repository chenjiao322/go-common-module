package utils

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "same", args: args{a: []string{"1", "2", "3", "4"}, b: []string{"1", "2", "3", "4"}}, want: nil},
		{name: "nil1", args: args{a: nil, b: []string{"1", "2", "3", "4"}}, want: nil},
		{name: "nil2", args: args{a: []string{"1", "2", "3", "4"}, b: nil}, want: []string{"1", "2", "3", "4"}},
		{name: "nil3", args: args{a: nil, b: nil}, want: nil},
		{name: "diff1", args: args{a: []string{"1", "2", "3", "4"}, b: []string{"1", "5"}}, want: []string{"2", "3", "4"}},
		{name: "diff2", args: args{a: []string{"1", "2"}, b: []string{"1", "2", "3"}}, want: nil},
		{name: "diff2", args: args{a: []string{"1", "2"}, b: []string{}}, want: []string{"1", "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniq(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"a", args{a: []string{"1", "2"}}, []string{"1", "2"}},
		{"a", args{a: []string{"1", "1"}}, []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uniq(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
