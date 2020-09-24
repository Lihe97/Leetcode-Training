package main

func findMode(root *TreeNode) []int {
	
	
	res := []int{}

	var base,count,maxCount int

	update := func(x int) {
		if x == base{
			count ++
		}else{
			base,count = x,1
		}
		if count == maxCount{
			res = append(res,base)
		}else if count > maxCount{
			maxCount = count
			res = []int{base}
		}

	}
	
	cur := root
	
	for cur != nil{
		if cur.Left == nil{
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur{
			pre = pre.Right
		}
		if pre.Right == nil{
			pre.Right = cur
			cur = cur.Left
		}else{
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func main() {
	
}
