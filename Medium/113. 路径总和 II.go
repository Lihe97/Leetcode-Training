package main

func pathSum(root *TreeNode, sum int) [][]int {

	res := [][]int{}
	temp := []int{}

	d(root ,sum,temp,&res)
	return res

}
func d(root *TreeNode,sum int,temp []int,res *[][]int){

	if root == nil{
		return
	}
	sum -= root.Val

	temp = append(temp,root.Val)

	if root.Left == nil && root.Right == nil && sum == 0{

		t := make([]int,len(temp))
		copy(t,temp)
		*res = append(*res,t)
	}
	d(root.Right,sum,temp,res)
	d(root.Left,sum,temp,res)
}

func main() {
	
}
