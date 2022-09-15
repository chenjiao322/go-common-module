package nacos_client

import (
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode("127.0.0.1", 4478, 100)
	node.UniqueKey()
	node.Ip()
	node.Port()
	node.Weight()
}

func TestNewNodes(t *testing.T) {
	nodes := NewNodes(NewRandomSelector())
	nodeList := []*Node{
		{"0.0.0.0", 1, 1},
		{"0.0.0.0", 2, 1},
		{"0.0.0.0", 3, 1},
		{"0.0.0.0", 4, 1},
		{"0.0.0.0", 5, 1},
	}
	nodes.Member()
	nodes.SetNode(nodeList...)
	fmt.Println(nodes.GetNode(RandomSelector, ""))
	fmt.Println(nodes.GetNode(HashSelector, "123"))
}
