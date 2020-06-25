package main

import "math"

var res int
func maxPathSum(root *TreeNode) int {
	res = math.MinInt8
	if root == nil{
		return res
	}
	dfs(root)
	return res
}
func dfs(root *TreeNode) int{
	if root == nil{
		return 0
	}
	left := max(0,dfs(root.Left))
	right := max(0,dfs(root.Right))
	res = max(res,root.Val + left +right)
	return root.Val + max(left,right)

}
func max(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

func main() {
	
}
