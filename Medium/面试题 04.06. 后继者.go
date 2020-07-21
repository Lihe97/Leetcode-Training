package main

import "fmt"

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	order := []*TreeNode{}

	queue := []*TreeNode{}

	for root != nil || len(queue) != 0{

		for root != nil{
			queue = append(queue,root)
			root = root.Left
		}
		root = queue[len(queue) - 1]
		order = append(order,root)
		queue = queue[0:len(queue) - 1]
		root = root.Right
	}


	for i := 0 ; i  < len(order) - 1; i ++{
		if order[i].Val == p.Val{
			return order[i+1]
		}
	}

	return nil

}

func main() {
	
	b := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	a := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: b,
	}
	fmt.Println(inorderSuccessor(a,a))

	
}
