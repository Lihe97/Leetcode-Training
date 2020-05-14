package main
func iisValidBST(root *TreeNode) bool {

	stack := []*TreeNode{}
	res := []int{}

	if root == nil{
		return true
	}

	for root != nil || len(stack) != 0{

		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if len(res) > 0 && root.Val <= res[len(res)-1]{
			return false
		}
		res = append(res,root.Val)
		root = root.Right

	}

	return true

}

func main() {
	
}
