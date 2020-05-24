package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {

	res := []string{}
	path := ""
	if root == nil{
		return res
	}

	run(&res,root,path)
	return res
}
func run(res *[]string,root *TreeNode,path string)  {

	if root.Left == nil && root.Right == nil{
		path += strconv.Itoa(root.Val)
		*res = append(*res,path)
		return
	}
	path += strconv.Itoa(root.Val)
	path += "->"
	if root.Left != nil{
		run(res,root.Left,path)
	}
	if root.Right != nil{
		run(res,root.Right,path)
	}
}

func main() {
	
}
