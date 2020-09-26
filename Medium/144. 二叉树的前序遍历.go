package main
//func preorderTraversal(root *TreeNode) []int {
//
//	res := []int{}
//	if root == nil{
//		return res
//	}
//
//	stack := []*TreeNode{}
//
//	stack = append(stack,root)
//
//	for len(stack) != 0{
//
//		temp := stack[len(stack) - 1]
//		res = append(res,temp.Val)
//		stack = stack[:len(stack) - 1]
//		if temp.Right != nil{
//			stack = append(stack,temp.Right)
//		}
//		if temp.Left != nil{
//			stack = append(stack,temp.Left)
//		}
//	}
//	return res
//}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}

	cur := root

	for cur != nil{
		if cur.Left == nil{
			res = append(res,cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur{
			pre = pre.Right
		}
		if pre.Right == nil{
			pre.Right = cur
			res = append(res,cur.Val)
			cur = cur.Left
		}else{
			pre.Right = nil
			cur = cur.Right
		}
	}
	return res
}
func main() {
	
}
