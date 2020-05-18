package main
func levelOrderBottom(root *TreeNode) [][]int {


	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue,root)
	for len(queue) != 0{
		temp := queue
		queue = []*TreeNode{}

		r := []int{}

		for len(temp) != 0{
			q := temp[0]
			temp = temp[1:]
			r = append(r,q.Val)
			if q.Left != nil{
				queue = append(queue,q.Left)
			}
			if q.Right != nil{
				queue = append(queue,q.Right)
			}
		}
		res = append(res,r)
	}
	for i := 0 ; i < len(res)/2 ; i ++{
		res[i],res[len(res)-i-1] = res[len(res)-i-1],res[i]
	}
	return res
}

func main() {
	
}
