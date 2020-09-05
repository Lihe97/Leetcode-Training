package main

import "fmt"

func getMaxLen(nums []int) int {

	res := -1

	stack := []int{}

	pre := -1

	for i := 0 ; i < len(nums) ; i ++{
		if nums[i] < 0 {
			stack = append(stack,i)
		}else if nums[i] == 0 {
			stack = []int{}
			pre = i
		}
		if len(stack) % 2 == 0{
			res = max(res,i-pre)
		}else{
			res = max(res,i-stack[0])
		}
	}
	return res
}
func max(a,b int) int {
	if a > b {
		return a
	}else{
		return b
	}
}



func main() {

	fmt.Println(getMaxLen([]int{-1,0,-1,1,1,-1,1,1}))






}