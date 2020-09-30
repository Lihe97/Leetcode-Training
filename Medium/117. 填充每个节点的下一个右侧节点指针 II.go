package main


type Node struct {
      Val int
      Left *Node
      Right *Node
      Next *Node
}

func connect(root *Node) *Node {

	if root == nil{
		return root
	}

	res := root


	queue := []*Node{}
	queue = append(queue,root)

	for len(queue) != 0{
		temp := queue
		queue = []*Node{}
		for i := 0 ; i < len(temp) -1; i ++{
			temp[i].Next = temp[i+1]
			if temp[i].Left != nil{
				queue = append(queue,temp[i].Left)
			}
			if temp[i].Right != nil{
				queue = append(queue,temp[i].Right)
			}
		}
		t := temp[len(temp)-1]
		if t.Left != nil{
			queue = append(queue,t.Left)
		}
		if t.Right != nil{
			queue = append(queue,t.Right)
		}
	}
	return res

}
func main() {
	
}
