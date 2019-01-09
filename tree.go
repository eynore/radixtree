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
			index += i
			if l > index {
				edge.node = &Node{edges: []*Edge{newEdge, &Edge{&Node{value, true, nil}, key[index:]}}}
			} else {
				edge.node = &Node{value, true, []*Edge{newEdge}}
			}
			return
		}
	}
	node.edges = append(node.edges, &Edge{&Node{value, true, nil}, key[index:]})

}

func (node *Node) leafNum() (n int) {
	if node.isLeaf() {
		return 0
	}

	for _, edge := range node.edges {
		if edge.node.isLeaf() {
			n++
		} else {
			n += edge.node.leafNum()
		}
	}
	return n
}

// ------------- String --------------

const labelWidth = 4
const nodeMargin = 2

type PNode struct {
	children []*Edge
	x        int
}

func (tree *Tree) String() (str string) {
	arr := []*PNode{&PNode{tree.root.edges, 0}}

	for len(arr) > 0 {
		var next []*PNode
		var line string
		for _, pnode := range arr {
			line += getN(pnode.x-len(line), " ")
			for _, edge := range pnode.children {
				node := edge.node
				x := len(line)
				if !node.isLeaf() {
					next = append(next, &PNode{node.edges, x})
				}
				width := getTreeWidth(node)
				line += getN(int(width/2)-labelWidth/2, " ")
				line += getLabel(edge.label)
				line += getN(x+width-len(line)+nodeMargin, " ")
			}
			line = line[:len(line)-2]
			line += "|"
		}

		str += line[:len(line)-1] + "\n"
		arr = next
	}
	return
}

func getN(n int, c string) (str string) {
	for i := 0; i < n; i++ {
		str += c
	}
	return
}
func getLabel(s string) string {
	l := len(s)
	if l > labelWidth {
		return s[:labelWidth]
	}
	return s + getN(labelWidth-l, " ")

}
func getTreeWidth(node *Node) int {
	n := node.leafNum()
	if n <= 1 {
		return labelWidth
	}
	return n*labelWidth + (n-1)*nodeMargin
}
