package main

func pathSum(root *TreeNode, sum int) int {

	if root == nil{
		return 0
	}

	res := help(root,sum)
	a := pathSum(root.Left,sum)
	b := pathSum(root.Right,sum)
	return res + a + b

}
func help(root *TreeNode,sum int) int {

	if root == nil{
		return 0
	}
	sum -= root.Val
	res := 0
	if  sum == 0{
		res ++
	}
	return res + help(root.Left,sum) + help(root.Right,sum)

}
func main() {
	
}
