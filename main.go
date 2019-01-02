// +build omit

package main

import (
	"fmt"

	"github.com/eynore/radixtree"
)

func main() {
	tree := radixtree.New()
	tree.Insert("tony", "haha")
	tree.Insert("tonyx", "hahax")
	tree.Insert("tonyy", "haha")

	value, ok := tree.Lookup("tony")

	if ok {
		fmt.Println(value.(string))
	} else {
		fmt.Println("not found")
	}
	value, ok = tree.Lookup("tonyx")

	if ok {
		fmt.Println(value.(string))
	} else {
		fmt.Println("not found")
	}
	value, ok = tree.Lookup("tonyy")

	if ok {
		fmt.Println(value.(string))
	} else {
		fmt.Println("not found")
	}

}
