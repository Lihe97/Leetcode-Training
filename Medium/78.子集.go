package main

import (
	"fmt"
)

//func subsets(nums []int) [][]int {
//
//	n  := len(nums)
//
//	res := [][]int{}
//
//	for i  := 0 ; i < 1<<n ; i ++{
//		temp := []int{}
//
//		for j := 0 ; j < len(nums) ; j ++{
//			if i >> j & 1 == 1{
//				temp = append(temp,nums[j])
//			}
//		}
//		res = append(res,temp)
//	}
//	return res
//}

func subsets(nums []int) [][]int {

	res := [][]int{}


	dfs(0,nums,&res,[]int{})

	return res
}
func dfs(cur int,nums []int,res *[][]int,temp []int)  {
	if cur == len(nums){
		*res = append(*res,append([]int{},temp...))
		return
	}
	temp = append(temp,nums[cur])
	dfs(cur + 1,nums,res,temp)
	temp = temp[:len(temp)-1]
	dfs(cur+1,nums,res,temp)
}


func main() {
	nums:=[]int{1,2,3,4,5,6}

	fmt.Println(subsets(nums))

}
