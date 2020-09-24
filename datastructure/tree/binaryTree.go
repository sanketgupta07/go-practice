package tree

import (
	"fmt"
	"io"
	"log"
)

/*
TreeNode ...
Simple Node
*/
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int64
}

/*
BinaryTree ...
Binary tree
*/
type BinaryTree struct {
	Root *TreeNode
}

/*
Insert ...
Insert a node in binary tree.
*/
func (t *BinaryTree) Insert(i int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &TreeNode{
			left:  nil,
			right: nil,
			value: i,
		}
	} else {
		t.Root.insertNode(i)
	}
	return t
}

func (node *TreeNode) insertNode(value int64) {
	if node == nil {
		log.Println("node is empty. can not add child to an null node")
		return
	} else if value <= node.value {
		if node.left == nil {
			node.left = &TreeNode{
				left:  nil,
				right: nil,
				value: value,
			}
		} else {
			node.left.insertNode(value)
		}
	} else {
		if node.right == nil {
			node.right = &TreeNode{
				left:  nil,
				right: nil,
				value: value,
			}
		} else {
			node.right.insertNode(value)
		}
	}
}

func PrintTree(w io.Writer, node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	PrintTree(w, node.left, ns+2, 'L')
	PrintTree(w, node.right, ns+2, 'R')
}

/*
PreOrder ...
read root
read left
read right
*/
func PreOrder() {

}

/*
InOrder ...
Read left
read root
read right
*/
func InOrder() {

}

/*
PostOrder ...
Read left
read right
read root
*/
func PostOrder() {

}
