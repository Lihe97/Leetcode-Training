package main

import "fmt"


func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return twonode(root.Left,root.Right)
}
func twonode(one *TreeNode,two *TreeNode)bool{
	if one == nil && two == nil{
		return true
	}
	if one == nil || two == nil{
		return false
	}
	if one.Val == two.Val {
		return twonode(one.Left,two.Right) && twonode(one.Right,two.Left)
	}else{
		return false
	}
}
func main() {
	g := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	f := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
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
	Left:  f,
	Right: g,
	}
	b := &TreeNode{
		Val:   3,
		Left:  d,
		Right: e,
	}
	a := &TreeNode{
		Val:   2,
		Left:  b,
		Right: c,
	}
	fmt.Println(isSymmetric(a))
	
}
