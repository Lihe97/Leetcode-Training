package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil{
		return 0
	}
	if root.Left == nil{
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil{
		return 1 + minDepth(root.Left)
	}

	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	e := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	d := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	c := &TreeNode{
		Val:   3,
		Left:  e,
		Right: d,
	}
	b := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	a := &TreeNode{
		Val:   2,
		Left:  b,
		Right: c,
	}
	fmt.Println(minDepth(a))
}
