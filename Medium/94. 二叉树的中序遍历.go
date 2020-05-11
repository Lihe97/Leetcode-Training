package main


//func inorderTraversal(root *TreeNode) []int {
//
//	res := []int{}
//
//	inorder(root,&res)
//	return res
//}
//func inorder(root *TreeNode,res *[]int){
//
//	if root.Left != nil{
//		inorder(root.Left,res)
//	}
//	*res = append(*res, root.Val)
//	if root.Right != nil{
//		inorder(root.Right,res)
//	}
//}
func inorderTraversal(t *TreeNode) []int {

	res := []int{}
	if t == nil{
		return res
	}

	stack := []*TreeNode{}

	for t != nil || len(stack) != 0{

		for t != nil{
			stack = append(stack,t)
			t = t.Left
		}

		t = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res,t.Val)
		t = t.Right

	}
	return res

}
func main() {
	
}
