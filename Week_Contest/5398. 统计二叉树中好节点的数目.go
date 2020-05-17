package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func goodNodes(root *TreeNode) int {
	res := 0
	//max := -math.MaxInt8
	find(root,&res,root.Val)
	return res
}
func find(root *TreeNode,res *int,p int){

	if root == nil{
		return
	}
	if p <= root.Val{
		 p = root.Val
		 *res ++
	}
	find(root.Left,res,p)
	find(root.Right,res,p)

}


func main() {

	fmt.Println(math.MaxInt8)
	fmt.Println(-2<<31)
	
}
