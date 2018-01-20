package main

/*
PROPERTIES:

1. Every node is either red or black.
2. The root is black.
3. Every leaf (NIL) is black.
4. If a node is red, then both its children are black.
5. For each node, all simple paths from the node to descendant leaves contain the
same number of black nodes
*/

import (
	"fmt"
	"strings"
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
	RedBlackTreeFixup(tree, &newNode)
}

func RedBlackTreeFixup(tree *RedBlackTree, newNode *Node) {
	// only properties 4. and 2. can be violated by an insertion
	iterator := newNode
	for iterator.parent.color == red {
		// when property 4. is violated
		if iterator.parent == iterator.parent.parent.left {
			// iterator's parent is the left child of its grandparent
			uncle := iterator.parent.parent.right
			if uncle.color == red { // case 1, only recoloring
				iterator.parent.color = black
				uncle.color = black
				iterator.parent.parent.color = red
				iterator = iterator.parent.parent
			} else {
				// uncle is black, needs rotations
				if iterator == iterator.parent.right {
					// case 2
					iterator = iterator.parent
					RedBlackTreeLeftRotate(tree, iterator)
				}
				// case 3
				iterator.parent.color = black
				iterator.parent.parent.color = red
				RedBlackTreeRightRotate(tree, iterator.parent.parent)
			}
		} else {
			// iterator's parent is the right child of its grandparent; symmetric
			uncle := iterator.parent.parent.left
			if uncle.color == red {
				iterator.parent.color = black
				uncle.color = black
				iterator.parent.parent.color = red
				iterator = iterator.parent.parent
			} else {
				if iterator == iterator.parent.left {
					iterator = iterator.parent
					RedBlackTreeRightRotate(tree, iterator)
				}
				iterator.parent.color = black
				iterator.parent.parent.color = red
				RedBlackTreeLeftRotate(tree, iterator.parent.parent)
			}
		}
	}
	// for property 2.
	tree.root.color = black
}

func RedBlackTreePrint(tree *RedBlackTree, node *Node, tabs int) {
	if node == tree.dummy {
		return
	}

	tabChars := strings.Repeat("\t", tabs)
	fmt.Print(tabChars)

	switch {
	case tree.root == node:
		fmt.Print("(root) ")
	case node == node.parent.left:
		fmt.Print("(left) ")
	case node == node.parent.right:
		fmt.Print("(right) ")
	}

	if node.color == red {
		fmt.Print("(red) ")
	} else {
		fmt.Print("(black) ")
	}
	fmt.Println(node.value)
	RedBlackTreePrint(tree, node.left, tabs+1)
	RedBlackTreePrint(tree, node.right, tabs+1)
}

func main() {
	var tree *RedBlackTree = RedBlackTreeCreate()
	RedBlackTreeInsert(tree, 11)
	RedBlackTreeInsert(tree, 2)
	RedBlackTreeInsert(tree, 14)
	RedBlackTreeInsert(tree, 1)
	RedBlackTreeInsert(tree, 7)
	RedBlackTreeInsert(tree, 5)
	RedBlackTreeInsert(tree, 8)
	RedBlackTreeInsert(tree, 15)
	RedBlackTreeInsert(tree, 4)
	RedBlackTreeInsert(tree, 3)
	RedBlackTreePrint(tree, tree.root, 0)
}
