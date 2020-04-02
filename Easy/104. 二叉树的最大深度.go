package main

func maxDepth(root *TreeNode) int {

	if root == nil{
		return 0
	}
	if root.Right == nil{
		return maxDepth(root.Left) + 1
	}
	if root.Left == nil{
		return maxDepth(root.Right) + 1
	}
	return max(maxDepth(root.Right),maxDepth(root.Left)) + 1

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	
}
