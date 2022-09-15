package nacos_client

import (
	"strconv"
	"sync"
)

type Node struct {
	ip     string
	port   int
	weight int
}

func NewNode(ip string, port int, weight int) *Node {
	return &Node{ip: ip, port: port, weight: weight}
}

func (n *Node) UniqueKey() string {
	return n.ip + ":" + strconv.Itoa(n.port)
}

func (n *Node) Ip() string {
	return n.ip
}

func (n *Node) Port() int {
	return n.port
}

func (n *Node) Weight() int {
	return n.weight
}

type Nodes struct {
	nodes     []*Node
	selectors map[SelectorName]Selector
	lock      sync.Mutex
}

func NewNodes(selectors ...Selector) *Nodes {
	s := make(map[SelectorName]Selector, 0)
	for _, v := range selectors {
		s[v.Name()] = v
	}
	return &Nodes{selectors: s, nodes: make([]*Node, 0)}
}

func (n *Nodes) SetNode(nodes ...*Node) {
	n.lock.Lock()
	defer n.lock.Unlock()

	n.nodes = nodes
	for _, selector := range n.selectors {
		selector.SetNodes(nodes...)
	}
}

func (n *Nodes) Member() []*Node {
	return n.nodes
}

func (n *Nodes) GetNode(selector SelectorName, hash string) (*Node, error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if s, ok := n.selectors[selector]; ok {
		return s.GetNode(n.nodes, hash)
	} else {
		return nil, NoSelectorError
	}
}
