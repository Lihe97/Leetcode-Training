package main


func rightSideView(root *TreeNode) []int {

	queue := []*TreeNode{}
	queue = append(queue, root)
	res := []int{}
	if root == nil{
		return res
	}
	for len(queue) != 0{
		temp := queue
		queue = []*TreeNode{}
		for i := 0 ; i < len(temp) ; i ++{
			if temp[i].Left != nil{
				queue = append(queue, temp[i].Left)
			}
			if temp[i].Right != nil{
				queue = append(queue, temp[i].Right)
			}
			if i == len(temp) - 1{
				res = append(res, temp[i].Val)
			}
		}
	}
	return res

}
func main() {
	
}
