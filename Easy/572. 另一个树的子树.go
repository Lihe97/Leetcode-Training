package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isSubtree(s *TreeNode, t *TreeNode) bool {


	if s == nil{
		return false
	}
	if t == nil{
		return true
	}

	return ss(s,t) || isSubtree(s.Left,t) || isSubtree(s.Right,t)
}
func ss(s *TreeNode, t *TreeNode) bool{

	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val{
		return false
	}
	return ss(s.Left, t.Left) && ss(s.Right, t.Right);


}

func main() {
	
}
