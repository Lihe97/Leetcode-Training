package main

import "fmt"

func permute(nums []int) [][]int {

	flag := make([]bool,len(nums))
	res := make([][]int,0)
	pth := make([]int,0)

	res = dfs(nums,&res,pth,flag)
	return res
}
func dfs(nums []int,res *[][]int,pth []int,flag []bool)[][]int{

	if len(pth) == len(nums){
		fmt.Println(res)
		*res = append(*res, pth)
	}else{
		for i := 0 ; i < len(flag); i ++ {
			if flag[i] == false {

				flag[i] = true
				dfs(nums, res, append(pth,nums[i]), flag)
				flag[i] = false
			}
		}
	}
	return *res

}

func main() {

	a := []int{1,2,3,4}
	fmt.Println(permute(a))
	
}
