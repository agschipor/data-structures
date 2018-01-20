package main

import (
	"fmt"
)

const (
	red   = iota
	black = iota
)

type Node struct {
	left   *Node
	parent *Node
	right  *Node
	color  int // red = 0; black = 1
	value  int
}

type RedBlackTree struct {
	dummy *Node
	root  *Node
}

func RedBlackTreeCreate() *RedBlackTree {
	var tree = RedBlackTree{}
	tree.dummy = &Node{}
	tree.dummy.color = black
	tree.dummy.parent, tree.dummy.left, tree.dummy.right = nil, nil, nil

	tree.root = tree.dummy

	return &tree
}

func RedBlackTreeLeftRotate(tree *RedBlackTree, x *Node) {
	y := x.right
	x.right = y.left
	if y.left != tree.dummy {
		y.left.parent = x
	}
	y.parent = x.parent
	switch {
	case x.parent == tree.dummy:
		tree.root = y
	case x == x.parent.left:
		x.parent.left = y
	case x == x.parent.right:
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func RedBlackTreeRightRotate(tree *RedBlackTree, y *Node) {
	x := y.left
	y.left = x.right
	if x.right != tree.dummy {
		x.right.parent = y
	}
	x.parent = y.parent
	switch {
	case y.parent == tree.dummy:
		tree.root = x
	case y == y.parent.left:
		y.parent.left = x
	case y == y.parent.right:
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

func RedBlackTreeInsert(tree *RedBlackTree, value int) {
	var newNode = Node{}
	newNode.color = red
	newNode.value = value

	parent := tree.dummy
	iterator := tree.root

	for iterator != tree.dummy {
		parent = iterator
		if newNode.value < iterator.value {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}

	newNode.parent = parent
	switch {
	case parent == tree.dummy:
		tree.root = &newNode
	case newNode.value < parent.value:
		parent.left = &newNode
	case newNode.value > parent.value: // or default
		parent.right = &newNode
	}

	newNode.left = tree.dummy
	newNode.right = tree.dummy
}

func main() {
	var tree *RedBlackTree = RedBlackTreeCreate()
	RedBlackTreeInsert(tree, 1)
	fmt.Println(tree, tree.root)
}
