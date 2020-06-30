package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func path(root *TreeNode) int{

	if root == nil{
		return 0
	}

	left := f(root,'L',0)

	right := f(root,'R',0)

	return max(left,right) - 1
}

func f(root *TreeNode,option byte,temp int) int {

	if root == nil{
		return temp
	}
	if option == 'L'{
		return f(root.Right,'R',temp+1)
	}else{
		return f(root.Left,'L',temp+1)
	}

}

func max(a,b int) int  {
	if a > b{
		return a
	}else{
		return b
	}
}




func main() {


}


