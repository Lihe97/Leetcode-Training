package main
func preorderTraversal(root *TreeNode) []int {

	res := []int{}
	if root == nil{
		return res
	}

	stack := []*TreeNode{}

	stack = append(stack,root)

	for len(stack) != 0{

		temp := stack[len(stack) - 1]
		res = append(res,temp.Val)
		stack = stack[:len(stack) - 1]
		if temp.Right != nil{
			stack = append(stack,temp.Right)
		}
		if temp.Left != nil{
			stack = append(stack,temp.Left)
		}
	}
	return res
}

func main() {
	
}
