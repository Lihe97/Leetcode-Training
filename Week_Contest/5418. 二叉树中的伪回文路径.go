package main

import "strconv"

func pseudoPalindromicPaths (root *TreeNode) int {
	temp := binaryTreePaths(root)
	res := 0

	for i := 0 ; i < len(temp) ; i ++{
		mp := map[byte]int{}
		ji := 0

		for j := 0 ; j < len(temp[i]) ; j++{
			mp[temp[i][j]] ++
		}
		for _,v := range mp{
			if v % 2 == 1{
				ji ++
			}
			if ji > 1{
				break
			}
		}
		if ji == 1{
			res ++
		}
	}
	return res

}
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string,0)
	s := ""
	if root == nil {
		return res
	}


	Run(&res,root,s)
	return res
}


func Run(res *[]string,root *TreeNode,s string) string {

	if root.Left == nil && root.Right == nil {
		s += strconv.Itoa(root.Val)
		*res = append(*res,s)
		return ""
	}

	s += strconv.Itoa(root.Val)
	if root.Left != nil {
		Run(res,root.Left,s)
	}
	if root.Right != nil {
		Run(res,root.Right,s)
	}
	return s
}


func main() {
	
}
