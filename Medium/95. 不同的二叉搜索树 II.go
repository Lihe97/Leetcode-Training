package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func generateTrees(n int) []*TreeNode {

	if n == 0{
		return nil
	}
	return h(1,n)

}
func h(left,right int) []*TreeNode {

	if left > right{
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}

	for i := left ; i <= right ; i ++{
		//fmt.Println("?")
		leftTrees := h(left,i-1)
		rightTrees := h(i+1,right)
		//fmt.Println(leftTrees,rightTrees)

		for _, a := range leftTrees{
			for _,b := range rightTrees{
				//fmt.Println("????????")
				cur := &TreeNode{
					Val:   i,
					Left:  nil,
					Right: nil,
				}
				cur.Left = a
				cur.Right = b
				//fmt.Println(left,right)
				res = append(res,cur)
			}
		}
	}
	return res
}


func main() {

	fmt.Println(generateTrees(3))
	
}
