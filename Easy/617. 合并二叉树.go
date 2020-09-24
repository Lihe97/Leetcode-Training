package main


func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil{
		return t2
	}
	if t2 == nil{
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left,t2.Left)
	t1.Right = mergeTrees(t1.Right,t2.Right)
	return t1
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil{
		return nil
	}

	queue := []*TreeNode{}
	if t1 == nil{
		return t2
	}
	if t2 == nil{
		return t1
	}
	queue = append(queue,t1)
	queue = append(queue,t2)

	for len(queue) > 0{
		temp1 := queue[0]

		temp2 := queue[1]

		temp1.Val += temp2.Val

		if temp1.Left != nil && temp2.Left != nil{
			queue = append(queue,temp1.Left)
			queue = append(queue,temp2.Left)
		}
		if temp1.Left == nil && temp2.Left != nil{
			temp1.Left = temp2.Left
		}
		if temp1.Right != nil && temp2.Right != nil{
			queue = append(queue,temp1.Right)
			queue = append(queue,temp2.Right)
		}

		if temp1.Right == nil && temp2.Right != nil{
			temp1.Right = temp2.Right
		}
		queue = queue[2:]
	}

	return t1
}
func main() {
	
}
