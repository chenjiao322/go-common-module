package utils

import (
	"testing"
)

func TestIsInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "999999"}, want: true},
		{name: "1", args: args{s: "9999999999999999999999999999999999999999999999999999"}, want: true},
		{name: "2", args: args{s: "1123.123"}, want: false},
		{name: "3", args: args{s: "a1123"}, want: false},
		{name: "4", args: args{s: " 1123"}, want: false},
		{name: "5", args: args{s: "1\r"}, want: false},
		{name: "6", args: args{s: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt(tt.args.s); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{a: []int{1, 2, 3, 4}}, want: 4},
		{name: "2", args: args{a: []int{1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{a: []int{1, 2, 3, 4}}, want: 1},
		{name: "2", args: args{a: []int{1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "1", args: args{16, 8}, want: 8},
		{name: "2", args: args{8, 16}, want: 8},
		{name: "3", args: args{30, 27}, want: 3},
		{name: "4", args: args{17, 1417241}, want: 1},
		{name: "4", args: args{8 * 17, 3 * 8}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickPow(t *testing.T) {
	type args struct {
		base uint64
		exp  uint64
		mod  uint32
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"1", args{base: 1, exp: 1, mod: 321}, 1},
		{"2", args{base: 0, exp: 0, mod: 123}, 1},
		{"3", args{base: 1283, exp: 12312, mod: 1241231}, 1006459},
		{"4", args{base: 1281233, exp: 14142312, mod: 178901}, 171066},
		{"5", args{base: 0, exp: 14142312, mod: 178901}, 0},
		{"6", args{base: 12381023, exp: 0, mod: 178901}, 1},
		{"7", args{base: 13446744073709551616, exp: 12446744073709551616, mod: 4194967296}, 2494358272},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.base, tt.args.exp, tt.args.mod); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
