package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)


	flag := make([]bool,len(nums))

	res := [][]int{}

	dfs(&res,nums,[]int{},flag)
	return res

}
func dfs(res *[][]int,nums []int,path []int,flag []bool)  {
	if len(path) == len(nums){
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0 ; i < len(nums); i ++{

		if flag[i] || i >= 1 && nums[i] == nums[i-1] && flag[i-1]{
			continue
		}else{
			path = append(path,nums[i])
			flag[i] = true
			dfs(res,nums,path,flag)
			flag[i] = false
			path = path[:len(path)-1]
		}
	}

}

func main() {

	fmt.Println(permuteUnique([]int{1,1,2}))
	
}
