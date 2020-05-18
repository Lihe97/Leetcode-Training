package main

import "math"

func isBalanced(root *TreeNode) bool {

	if root == nil{
		return true
	}
	return math.Abs(float64(depth(root.Left)) - float64(depth(root.Right))) < 2 && isBalanced(root.Left) && isBalanced(root.Right)


}
func depth(root *TreeNode) int{
	if root == nil{
		return 0
	}
	if root == nil{
		return depth(root.Right) + 1
	}
	if root == nil{
		return depth(root.Left) + 1
	}
	return max(depth(root.Left),depth(root.Right)) + 1
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}

func main() {
	
}
