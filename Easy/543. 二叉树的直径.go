package main

//type TreeNode struct {
//      Val int
//      Left *TreeNode
//      Right *TreeNode
//}
//求每个节点的左+右深度和，对每个节点进行比较
var max int = 0
func diameterOfBinaryTree(root *TreeNode) int {
	max=0
	high(root)
	return max
}
func high(root *TreeNode)int{
	if root==nil {
		return 0
	}
	left := high(root.Left)
	right := high(root.Right)
	if left+right>max{
		max=left+right
	}
	if left>right{
		return left+1
	}else {
		return right+1
	}
}
func main() {
	
}
