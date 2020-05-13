package main

import "fmt"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {

	res := [][]int{}
	if root == nil{
		return res
	}

	queue := []*TreeNode{}
	queue = append(queue,root)

	for len(queue) != 0{
		temp := []int{}
		q := queue
		queue = []*TreeNode{}

		for len(q) != 0{
			a := q[0]
			if a.Left != nil{
				queue = append(queue,a.Left)
			}
			if a.Right != nil{
				queue = append(queue,a.Right)
			}
			temp = append(temp,a.Val)
			q = q[1:]
		}
		res = append(res,temp)
		fmt.Println(res)
	}
	return res
}

func main() {
	e := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	d := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	c := &TreeNode{
		Val:   4,
		Left:  d,
		Right: e,
	}

	b := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	a := &TreeNode{
		Val:   2,
		Left:  b,
		Right: c,
	}
	fmt.Println(levelOrder(a))
	
}
