package main

func isSameTree(p *TreeNode, q *TreeNode) bool {



	if q == nil && p == nil{
		return true
	}
	if p == nil || q == nil{
		return false
	}
	if q.Val != p.Val{
		return false
	}
	return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}

func main() {

	c := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	b := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	a := &TreeNode{
		Val:   2,
		Left:  b,
		Right: c,
	}
}
