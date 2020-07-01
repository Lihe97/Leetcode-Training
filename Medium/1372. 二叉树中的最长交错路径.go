package main
//var res int = 0
func longestZigZag(root *TreeNode) int {

	res := 0

	f(root.Left,-1,1,&res)
	f(root.Right,1,1,&res)
	return res
}
func f(root *TreeNode,last int,cur int,res *int) {

	if root == nil{
		return
	}
	*res = max(*res,cur)
	if last == -1{
		f(root.Left,-1,1,res)
		f(root.Right,1,cur+1,res)
	}else{
		f(root.Left,-1,cur+1,res)
		f(root.Right,1,1,res)
	}
}
func max(a,b int) int{
	if a > b{
		return a
	}else {
		return b
	}
}

func main() {
	
}
