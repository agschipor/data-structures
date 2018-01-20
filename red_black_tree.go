package main

import (
	"fmt"
)

type Node struct {
	left   *Node
	parent *Node
	right  *Node
	isnil  bool
	value  int
}

func CreateHeadNode() *Node {
	var head Node = Node{}
	head.left = &head
	head.right = &head
	head.parent = &head
	head.isnil = true
	head.value = 0
	return &head
}

func main() {
	var head *Node = CreateHeadNode()
	fmt.Println(head.value, head, head.parent)
}
