package main
type Node struct {
	Val int
	Children []*Node
}


func llevelOrder(root *Node) [][]int {

	res := [][]int{}
	if root == nil{
		return res
	}

	queue := []*Node{}
	queue = append(queue,root)

	for len(queue) != 0{
		temp := []int{}
		q := queue
		queue = []*Node{}

		for len(q) != 0{
			a := q[0]
			for i := 0 ; i  < len(a.Children) ; i++{
				queue = append(queue,a.Children[i])
			}
			q = q[1:]
			temp = append(temp,a.Val)
		}
		res = append(res,temp)
	}
	return res
}

func main() {
	
}
