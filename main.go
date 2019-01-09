// +build omit

package main

import (
	"fmt"

	"github.com/eynore/radixtree"
)

func main() {
	tree := radixtree.New()

	tests := []struct {
		key   string
		value string
	}{
		{"tony", "1"},
		{"tonyx", "2"},
		{"tonyxx", "3"},
		{"tonyxy", "4"},
		{"to", "5"},
		{"tox", "6"},
		{"toy", "7"},
		{"xoy", "8"},
		{"abc", "9"},
		{"abd", "10"},
		{"abdc", "10"},
	}

	for _, test := range tests {
		tree.Insert(test.key, test.value)
	}

	fmt.Println(tree)

	fmt.Printf("%3s", "abcdef")
}
