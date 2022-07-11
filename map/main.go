package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	Name     string
	Children []*Node
}

type MapItem struct {
	A int
}

func main() {
	m := make(map[string]int)
	m["asd"] = 5
	delete(m, "asd")
	m["asd"] = 2
	fmt.Println(m)

	m["b"] = 10

	fmt.Println(m["asd"], m["c"])
	if value, found := m["b"]; found {
		fmt.Println("found", value)
	}

	nodes := make(map[string]Node)
	nodes["asd"] = Node{
		Children: make([]*Node, 0),
	}

	// for true {
	// 	var id string
	// 	if _, found := nodes[id]; !found {
	// 		nodes[id] = Node{
	// 			Name: "",
	// 			Children: make([]*Node, 0),
	// 		}
	// 	}
	// 	nodes[id].Children = append(nodes[id].Children, ...)
	// }

	items := make(map[int]*MapItem, 0)

	items[1] = &MapItem{
		A: 1,
	}
	fmt.Println(items)
	items[1].A = 2
	fmt.Println()
	spew.Dump(items)
	fmt.Println()
	fmt.Println(items)
}
