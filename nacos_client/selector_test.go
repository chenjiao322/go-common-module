package nacos_client

import (
	"fmt"
	"testing"
)

func TestNewRandomSelector(t *testing.T) {
	h := NewRandomSelector()
	nodeList := []*Node{
		{"0.0.0.0", 1, 1},
		{"0.0.0.0", 2, 1},
		{"0.0.0.0", 3, 1},
		{"0.0.0.0", 4, 1},
		{"0.0.0.0", 5, 1},
	}
	h.SetNodes(nodeList...)
	fmt.Println(h.GetNode(nodeList, "3"))
	fmt.Println(h.GetNode(nodeList, "4"))

	h.SetNodes()
	h.Name()
	fmt.Println(h.GetNode(nodeList, "3"))
}

func TestNewHashSelector(t *testing.T) {
	h := NewHashSelector()
	nodeList := []*Node{
		{"0.0.0.0", 1, 1},
		{"0.0.0.0", 2, 1},
		{"0.0.0.0", 3, 1},
		{"0.0.0.0", 4, 1},
		{"0.0.0.0", 5, 1},
	}
	h.Name()
	h.SetNodes(nodeList...)
	fmt.Println(h.GetNode(nodeList, "3"))
	fmt.Println(h.GetNode(nodeList, "4"))

}
