package data_structure

import (
	"strings"
	"sync"
)

// 用于uri的匹配, 所以采用了[]string作为参数
// * 表示一级通配符, \* 表示正常的*号
const wildcard = "*"
const star = `\*`
const realStar = "*"

type trieNode struct {
	children map[string]*trieNode
	WildCard *trieNode
	isWord   bool
}

type Trie struct {
	root *trieNode
	size int
	lock sync.RWMutex
}

func newTrieNode() *trieNode {
	return &trieNode{children: make(map[string]*trieNode, 0)}
}

func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

func (t *trieNode) GetWithDefault(s string) *trieNode {
	if s == wildcard {
		if t.WildCard == nil {
			t.WildCard = newTrieNode()
		}
		return t.WildCard
	}
	s = strings.ReplaceAll(s, star, realStar)
	if _, ok := t.children[s]; !ok {
		t.children[s] = newTrieNode()
	}
	return t.children[s]
}

func (t *Trie) Add(ss ...[]string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, s := range ss {
		node := t.root
		for _, v := range s {
			node = node.GetWithDefault(v)
		}
		if !node.isWord {
			t.size += 1
		}
		node.isWord = true
	}
}

type traceNodePair struct {
	trace  []string
	node   *trieNode
	isWord bool
}

// LongestPrefix 返回存在与trie内的最长前缀
// 如果存在多条符合匹配条件的前缀, 会返回通配符在比较后面的那一条
func (t *Trie) LongestPrefix(s []string) []string {
	t.lock.RLock()
	defer t.lock.RUnlock()
	var ans []string
	nodes := []*traceNodePair{{trace: []string{}, node: t.root}}
	for _, v := range s {
		nextNodes := make([]*traceNodePair, 0)
		for _, pair := range nodes {
			if next, ok := pair.node.children[v]; ok {
				nextNodes = append(nextNodes, &traceNodePair{
					trace:  copyAndAdd(pair.trace, v),
					node:   next,
					isWord: next.isWord,
				})
			}
			if pair.node.WildCard != nil {
				nextNodes = append(nextNodes, &traceNodePair{
					trace:  copyAndAdd(pair.trace, wildcard),
					node:   pair.node.WildCard,
					isWord: pair.node.WildCard.isWord,
				})
			}
		}
		nodes = nextNodes
		for _, pair := range nodes {
			if pair.isWord && len(pair.trace) > len(ans) {
				ans = pair.trace
			}
		}
	}
	return ans
}

func (t *Trie) Size() int {
	return t.size
}

func copyAndAdd(old []string, add string) []string {
	newSlice := make([]string, len(old)+1)
	copy(newSlice, old)
	newSlice[len(newSlice)-1] = add
	return newSlice
}
