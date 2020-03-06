package main
 type TreeNode struct {
 	     Val int
    Left *TreeNode
      Right *TreeNode
 }
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	q:= []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 && q[0] != nil{
		q = append(q, q[0].Left,q[0].Right)
		q = q[1:]
	}
	for  _,v := range q{
		if v != nil{
			return false
		}
	}
	return true

}

func main() {
	
}
