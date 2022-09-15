package nacos_client

import (
	"math/rand"
	"stathat.com/c/consistent"
)

type SelectorName string

const (
	RandomSelector = SelectorName("RandomSelector")
	HashSelector   = SelectorName("HashSelector")
)

type Selector interface {
	SetNodes(nodes ...*Node)
	GetNode(nodes []*Node, hash string) (*Node, error)
	Name() SelectorName
}

type randomSelector struct{}

func NewRandomSelector() *randomSelector {
	return &randomSelector{}
}

func (r *randomSelector) GetNode(nodes []*Node, _ string) (*Node, error) {
	if len(nodes) == 0 {
		return nil, NoAvailableNodeError
	}
	return nodes[rand.Intn(len(nodes))], nil
}

func (r *randomSelector) SetNodes(_ ...*Node) {
	// 不需要更新selector
}

func (r *randomSelector) Name() SelectorName { return RandomSelector }

type hashSelector struct {
	ring *consistent.Consistent
}

func NewHashSelector() *hashSelector {
	return &hashSelector{ring: consistent.New()}
}

func (h *hashSelector) SetNodes(nodes ...*Node) {
	h.ring = consistent.New()
	for _, node := range nodes {
		h.ring.Add(node.UniqueKey())
	}
}

func (h *hashSelector) GetNode(nodes []*Node, hash string) (*Node, error) {
	name, err := h.ring.Get(hash)
	if err != nil {
		return nil, NoAvailableNodeError
	}
	for _, node := range nodes {
		if node.UniqueKey() == name {
			return node, nil
		}
	}
	return nil, NoAvailableNodeError
}

func (h *hashSelector) Name() SelectorName {
	return HashSelector
}
