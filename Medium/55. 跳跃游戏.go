package main

import "fmt"

func canJump(nums []int) bool {
	c := make([]int,len(nums))
	for i := 0 ;i < len(nums);i++{
		c[i] = i + nums[i]
	}
	max := c[0]
	cur := 0
	for cur < len(nums) && cur <= max{
		if max < c[cur]{
			max = c[cur]
		}
		cur ++
	}
	if cur == len(nums){
		return true
	}
	return false



}
func main() {
	a := []int{2,3,1,1,4}
	fmt.Println(canJump(a))
}
