package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	res := &TreeNode{}
	pathP := getPath(root, p)
	pathQ := getPath(root, q)

	for i := 0 ; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i] ; i ++{
		res = pathQ[i]
	}
	return res
}
func getPath(root, target *TreeNode) (path []*TreeNode) {
	node := root
	for node != target {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}


func main() {
	
}
