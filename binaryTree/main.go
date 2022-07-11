package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	Value    *bool
	Operator *string
	Left     *Node
	Right    *Node
}

type InfiniteNode struct {
	Name     string
	Parent   *InfiniteNode
	Children []*InfiniteNode
}

func (n *Node) Eval() bool {
	if n.Value != nil {
		return *n.Value
	}
	switch *n.Operator {
	case "AND":
		return n.Left.Eval() && n.Right.Eval()
	case "OR":
		return n.Left.Eval() || n.Right.Eval()
	}
	return false //TODO: error handling
}

func main() {
	data, err := os.ReadFile("2.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	root := Node{}
	json.Unmarshal(data, &root)
	fmt.Println(root)
	spew.Dump(root)
	fmt.Println(root.Eval())
}
