package data_structure

import (
	"reflect"
	"testing"
)

func TestLongestPrefix(t1 *testing.T) {
	t := NewTrie()
	t.Add([]string{"a", "b"}, []string{"a", "b", "d"})
	if reflect.DeepEqual(t.LongestPrefix([]string{"a", "b", "e"}), []string{"a", "b"}) != true {
		t1.Errorf("LongestPrefix : %+v", t.LongestPrefix([]string{"a", "b", "e"}))
	}
}

func TestCornerCase(t1 *testing.T) {
	t := NewTrie()
	t.Add(nil, nil)
	t = NewTrie()
	t.Add([]string{"你"}, []string{"🎈😄"})
	if t.Size() != 2 {
		t1.Errorf("LongestPrefix: %+v", t.LongestPrefix([]string{"你", "👌"}))
	}

	if reflect.DeepEqual(t.LongestPrefix([]string{"你", "👌"}), []string{"你"}) != true {
		t1.Errorf("LongestPrefix: %+v", t.LongestPrefix([]string{"你", "👌"}))
	}
	if reflect.DeepEqual(t.LongestPrefix([]string{"你"}), []string{"你"}) != true {
		t1.Errorf("LongestPrefix: %+v", t.LongestPrefix([]string{"你"}))
	}
	if reflect.DeepEqual(t.LongestPrefix([]string{"kk"}), []string(nil)) != true {
		t1.Errorf("LongestPrefix: %+v", t.LongestPrefix([]string{"kk"}))
	}
}

func TrieForLongestPrefix() *Trie {
	ans := NewTrie()
	ans.Add([]string{"a", "*", "b"})
	ans.Add([]string{"*"})
	ans.Add([]string{"*", "b", "b"})
	ans.Add([]string{"a", "b", "c"})
	ans.Add([]string{"a", "b"})
	ans.Add([]string{"c", "d", "*"})
	ans.Add([]string{"c", "d", "*", "*"})
	return ans
}

func TestTrie_LongestPrefixWildCard(t1 *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		trie *Trie
		args args
		want []string
	}{
		{name: "1", trie: TrieForLongestPrefix(), args: args{s: []string{"a", "c", "c"}}, want: []string{"*"}},
		{name: "2", trie: TrieForLongestPrefix(), args: args{s: []string{"z", "b", "b"}}, want: []string{"*", "b", "b"}},
		{name: "3", trie: TrieForLongestPrefix(), args: args{s: []string{"b", "c", "c", "d"}}, want: []string{"*"}},
		{name: "4", trie: TrieForLongestPrefix(), args: args{s: []string{"a", "*", "b"}}, want: []string{"a", "*", "b"}},
		{name: "5", trie: TrieForLongestPrefix(), args: args{s: []string{"a", "b", "d"}}, want: []string{"a", "b"}},
		{name: "6", trie: TrieForLongestPrefix(), args: args{s: []string{"a", "c", "d"}}, want: []string{"*"}},
		{name: "7", trie: TrieForLongestPrefix(), args: args{s: []string{"a", "b", "c", "d"}}, want: []string{"a", "b", "c"}},
		{name: "8", trie: TrieForLongestPrefix(), args: args{s: []string{"c", "d", "e", "e"}}, want: []string{"c", "d", "*", "*"}},
		{name: "9", trie: TrieForLongestPrefix(), args: args{s: []string{"c", "d", "e"}}, want: []string{"c", "d", "*"}},
		{name: "10", trie: TrieForLongestPrefix(), args: args{s: []string{"c", "d"}}, want: []string{"*"}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.trie.LongestPrefix(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("LongestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
