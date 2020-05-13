package main

import "fmt"


func isValidBST(root *TreeNode) bool {

	if root == nil{
		return true
	}
	//var temp []int
	//add(root,&temp)
	//if len(temp) == 1{
	//	return true
	//for i := 1 ; i < len(temp) ; i ++{
	//	if temp[i] <= temp[i-1]{
	//		return false
	//	}
	//}
	t(root,)
	return true
}
//func add(root *TreeNode, temp *[]int){
//	if root.Left != nil{
//		add(root.Left,temp)
//	}
//	*temp = append(*temp,root.Val)
//	if root.Right != nil{
//		add(root.Right,temp)
//	}
//}
func t(root *TreeNode,last int) bool {

	if root.Left != nil{
		t(root.Left,last)
	}
}

func main() {
	
	e := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	
	d := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	
	c := &TreeNode{
		Val:   4,
		Left:  d,
		Right: e,
	}

	b := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	
	a := &TreeNode{
		Val:   2,
		Left:  b,
		Right: c,
	}
	fmt.Println(isValidBST(a))
	
	
}
