package main

import "fmt"


func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil{
		return 0
	}
	res := 0
	sum(root,&res)
	return res
}
func sum(root *TreeNode,res *int){
	if root.Left!= nil && root.Left.Left == nil{
		*res += root.Left.Val
	}
	if root.Left!= nil{
		sum(root.Left,res)
	}
	if root.Right!= nil{
		sum(root.Right,res)
	}
	
}


func main() {
	
	e := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	
	d := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	
	c := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	b := &TreeNode{
		Val:   2,
		Left:  d,
		Right: e,
	}
	
	a := &TreeNode{
		Val:   1,
		Left:  b,
		Right: c,
	}
	fmt.Println(sumOfLeftLeaves(a))
	
}
