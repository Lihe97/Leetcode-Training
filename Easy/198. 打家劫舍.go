package main

import "fmt"

func rob(nums []int) int {
	dp := make([]int,len(nums))
	if len(nums) == 0{
		return 0
	}
	dp[0] = nums[0]
	if len(nums) == 1{
		return dp[0]
	}
	if len(nums) == 2{
		return maxx(nums[0],nums[1])
	}
	if len(nums) == 3{
		return maxx(nums[0]+nums[2],nums[1])
	}
	dp[1] = nums[1]
	dp[2] = nums[2] + nums[0]
	for i:= 3;i < len(nums) ; i++{
		dp[i] = maxx(dp[i-3]+nums[i],dp[i-2]+nums[i])
	}
	//fmt.Println(dp)
	return maxx(dp[len(nums)-1],dp[len(nums)-2])


}
func maxx(a,b int)int{
	if a >  b{
		return a
	}else{
		return b
	}

}
func main() {
	a := []int{2,7,9,3,1}
	fmt.Println(rob(a))
}
