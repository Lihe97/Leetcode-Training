package main
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}

	if root == nil{
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue,root)

	for len(queue) != 0{
		var sum float64
		num := 0

		q := queue
		queue = []*TreeNode{}
		for len(q) != 0{
			a := q[0]
			q = q[1:]
			sum += float64(a.Val)
			num ++
			if a.Left != nil{
				queue = append(queue,a.Left)
			}
			if a.Right != nil{
				queue = append(queue,a.Right)
			}
		}
		res = append(res,sum/float64(num))
	}
	return res

}

func main() {
	
}
