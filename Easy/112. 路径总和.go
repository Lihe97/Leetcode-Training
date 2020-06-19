package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//func hasPathSum(root *TreeNode, sum int) bool {
//	if root == nil{
//		return false
//	}
//	temp := []int{}
//
//	b(&temp,root,0)
//	for i := 0 ; i < len(temp) ; i ++{
//		if temp[i] == sum{
//			return true
//		}
//	}
//	return false
//}
//func b(temp *[]int,root *TreeNode,sum int){
//
//	sum += root.Val
//	if root.Right == nil && root.Left == nil{
//		*temp = append(*temp,sum)
//		return
//	}
//	if root.Left != nil{
//		b(temp,root.Left,sum)
//	}
//	if root.Right != nil{
//		b(temp,root.Right,sum)
//
//	}
//
//}
func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil{
		return false
	}
	sum -= root.Val
	if root.Right == nil && root.Left == nil{
		return sum == 0
	}
	return hasPathSum(root.Left,sum) || hasPathSum(root.Right,sum)
}

func main() {
	
}
