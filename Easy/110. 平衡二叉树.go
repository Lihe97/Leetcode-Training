package main

import "math"

func isBalanced(root *TreeNode) bool {

	if root == nil{
		return true
	}

	return math.Abs(float64(depth(root.Left))-float64(depth(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}
func depth(root *TreeNode) int{
	if root == nil{
		return 0
	}
	if root.Left == nil{
		return depth(root.Right) + 1
	}
	if root.Right == nil{
		return depth(root.Left) + 1
	}
	return max(depth(root.Left),depth(root.Right)) + 1
}
func max(a,b int)int {
	if a > b{
		return a
	}else{
		return b
	}
}

func main() {
	
}
