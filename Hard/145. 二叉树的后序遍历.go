package main

import "fmt"

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}
//递归
//func postorderTraversal(root *TreeNode) []int {
//
//	res := []int{}
//
//	f(root,&res)
//
//	return res
//}
//func f(root *TreeNode,res *[]int){
//	if root == nil{
//		return
//	}
//
//	f(root.Left,res)
//	f(root.Right,res)
//	*res = append(*res,root.Val)
//}

//非递归
//func postorderTraversal(root *TreeNode) []int {
//	res := []int{}
//	if root == nil{
//		return res
//	}
//
//	stack1 := []*TreeNode{}
//	stack2 := []*TreeNode{}
//
//	stack1 = append(stack1,root)
//
//	for len(stack1) != 0{
//		temp := stack1[len(stack1) - 1]
//		stack2 = append(stack2,temp)
//		if temp.Left != nil{
//			stack1 = append(stack1,temp.Left)
//		}
//		if temp.Right != nil{
//			stack1 = append(stack1,temp.Right)
//		}
//	}
//	for len(stack2) != 0{
//		res = append(res,stack2[len(stack2)-1].Val)
//		stack2 = stack2[:len(stack2)-1]
//
//	}
//	return res
//}


//一个栈
func postorderTraversal(root *TreeNode) []int {

	stack := []*TreeNode{}
	stack = append(stack,root)
	res := []int{}

	for len(stack) != 0{

		temp := stack[len(stack) - 1]
		if temp == nil{
			stack = stack[:len(stack) - 1]
			temp = stack[len(stack) - 1]
			res = append(res,temp.Val)
			stack = stack[:len(stack) - 1]
			continue

		}
		stack = append(stack,nil)

		if temp.Right != nil{
			stack = append(stack,temp.Right)
		}
		if temp.Left != nil{
			stack = append(stack,temp.Left)
		}
	}
	return res
}
func main() {

	c := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	b := &TreeNode{
		Val:   2,
		Left:  c,
		Right: nil,
	}
	a := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: b,
	}
	fmt.Println(postorderTraversal(a))
}
