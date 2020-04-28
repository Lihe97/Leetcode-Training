package main

func convertBST(root *TreeNode) *TreeNode {
	var sum = 0
	dfss(root,&sum)
	return root

}
func dfss(root *TreeNode,sum *int){
	if root == nil{
		return
	}

	dfss(root.Right,sum)
	*sum += root.Val
	root.Val = *sum
	dfss(root.Left,sum)
}

func main() {


	
}
