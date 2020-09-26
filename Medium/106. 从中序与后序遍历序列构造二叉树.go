package main

import "fmt"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	for k := range inorder{
		if inorder[k] == postorder[len(postorder)-1]{
			return &TreeNode{
				Val:postorder[len(postorder)-1],
				Left:buildTree(inorder[:k],postorder[:k]),
				Right:buildTree(inorder[k+1:],postorder[k:len(postorder)-1]),
			}
		}
	}
	return nil
}
//func helper(start,end int,mp map[int]int,pos []int) *TreeNode {
//	if start > end{
//		return nil
//	}
//	if len(pos) == 0{
//		return nil
//	}
//	//fmt.Println(pos)
//	val := pos[len(pos)-1]
//	root := &TreeNode{Val:val}
//	pos = pos[:len(pos)-1]
//
//	p := mp[val]
//
//
//	root.Right = helper(p+1,end,mp,pos)
//	root.Left = helper(start,p-1,mp,pos)
//	return root
//}

func main() {

	fmt.Println(buildTree([]int{9,3,15,20,7},[]int{9,15,7,20,3}))
	
}
