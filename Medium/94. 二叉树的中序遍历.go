package main


func inorderTraversal(root *TreeNode) []int {

	res := []int{}

	inorder(root,&res)
	return res
}
func inorder(root *TreeNode,res *[]int){

	if root.Left != nil{
		inorder(root.Left,res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil{
		inorder(root.Right,res)
	}
}
func main() {
	
}
