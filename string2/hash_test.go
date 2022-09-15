package string2

import (
	"reflect"
	"testing"
)

func assert(ok bool) {
	if !ok {
		panic("assert fail")
	}
}

func TestRabinKarp1(t *testing.T) {
	match := 0
	pattern := "vklqa"
	target := NewRabinKarp().SetUp([]byte(pattern)).Hash()
	text := "agfwenfvklqanfc"
	for _, v := range Scan([]byte(text), len(pattern)) {
		if v.Hash == target {
			assert(text[v.Left:v.Right] == pattern)
			match += 1
		}
	}
	assert(match == 1)
}

func TestRabinKarp2(t *testing.T) {
	match := 0
	pattern := "fqajdef"
	text := "agfwenffvmwiosvwnjqegvwaejgfqajdefadoalkdascmlkascmfqajdefvklqanfc"
	target := NewRabinKarp().SetUp([]byte(pattern)).Hash()
	for _, v := range Scan([]byte(text), len(pattern)) {
		if v.Hash == target {
			assert(text[v.Left:v.Right] == pattern)
			match += 1
		}
	}
	assert(match == 2)
}

func Test_rabinKarp_SetUp(t *testing.T) {
	tests := []struct {
		name string
		args string
		want uint64
	}{
		{name: "1", args: "abcde", want: 2923770999},
		{name: "2", args: "vwegvwgvwsgswt", want: 281296374},
		{name: "3", args: "fcjwnebgvjwnbvnalovnakivnolanv", want: 2392766693},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRabinKarp()
			if got := r.SetUp([]byte(tt.args)); !reflect.DeepEqual(got.Hash(), tt.want) {
				t.Errorf("SetUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rabinKarp2(t *testing.T) {
	rk := NewRabinKarp().SetUp([]byte("fcjwnebgvjwnbvnalovnakivnolanv"))
	assert(rk.PopLeft().Hash() == 3895341290)
	assert(rk.PopLeft().Hash() == 3545919536)
	assert(rk.PopLeft().Hash() == 3436900152)
	assert(rk.PopLeft().Hash() == 746204803)
	assert(rk.PopLeft().Hash() == 1344168003)
	assert(rk.PopLeft().Hash() == 3829350982)
	assert(rk.PopLeft().Hash() == 1247435130)
}

func Test_rabinKarpForCoverage(t *testing.T) {
	assert(len(Scan([]byte("123123"), 999)) == 0)
	assert(string(NewRabinKarp().SetUp([]byte("123")).Bytes()) == "123")
	defer func() { recover() }()
	NewRabinKarp().PopLeft()
}
