package main
func isSymmetric(root *TreeNode) bool {

	if root == nil{
		return true
	}

	return isSame(root.Left,root.Right)
}
func isSame(l,r *TreeNode) bool {

	if l == nil && r == nil{
		return true
	}
	if l == nil || r == nil{
		return false
	}
	if l.Val == r.Val{
		return isSame(l.Left,r.Right) && isSame(l.Right,r.Left)
	}else{
		return false
	}

}

func main() {
	
}
