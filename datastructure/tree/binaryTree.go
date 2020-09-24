package tree

import (
	"fmt"
	"io"
	"log"
)

/*
Node ...
Simple Node
*/
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int64
}

type BinaryTree struct {
	root *TreeNode
}

/*
Insert ...
Insert a node in binary tree.
*/
func (t *BinaryTree) Insert(i int64) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{
			left:  nil,
			right: nil,
			value: i,
		}
	} else {
		t.root.insertNode(i)
	}
	return t
}

func (node *TreeNode) insertNode(value int64) {
	if node == nil {
		log.Println("node is empty. can not add child to an null node")
		return
	} else if node.value <= value {
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

func Print(w io.Writer, node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
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
