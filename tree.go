package radixtree

import (
	"strings"
)

type Node struct {
	value    interface{}
	hasValue bool
	edges    []*Edge
}

func (node *Node) isLeaf() bool {
	return len(node.edges) == 0
}

type Edge struct {
	node  *Node
	label string
}

type Tree struct {
	root *Node
}

func New() *Tree {
	return &Tree{&Node{}}
}

func (tree *Tree) Lookup(key string) (value interface{}, ok bool) {
	if node, ok, _ := tree.lookup(key); ok {
		return node.value, node.hasValue
	}
	return
}

func (tree *Tree) lookup(key string) (node *Node, ok bool, index int) {
	node = tree.root
	l := len(key)

loop:
	for index < l && !node.isLeaf() {
		for _, edge := range node.edges {
			if strings.HasPrefix(key[index:], edge.label) {
				node = edge.node
				index += len(edge.label)
				continue loop
			}
		}
		break
	}
	return node, index == l && node != tree.root, index
}

func (tree *Tree) Insert(key string, value interface{}) {
	node, ok, index := tree.lookup(key)
	if ok {
		node.value = value
		node.hasValue = true
		return
	}
	l := len(key) - index
	for _, edge := range node.edges {
		label := edge.label
		ll := len(label)
		if ll > l {
			ll = l
		}
		i := 1
		for ; i <= ll; i++ {
			if strings.HasPrefix(key[index:], label[:i]) {
				continue
			}
			break
		}
		i--
		if i > 0 {
			edge.label = label[:i]
			newEdge := &Edge{edge.node, label[i:]}
			if l > index+i {
				edge.node = &Node{edges: []*Edge{newEdge, &Edge{&Node{value, true, nil}, key[index+i:]}}}
			} else {
				edge.node = &Node{value, true, []*Edge{newEdge}}
			}
			return
		}
	}
	node.edges = append(node.edges, &Edge{&Node{value, true, nil}, key[index:]})

}
