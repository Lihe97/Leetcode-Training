package main

import "fmt"

func constrainedSubsetSum(nums []int, k int) int {

	dp := make([]int,len(nums))

	queue := []int{}

	dp[0] = nums[0]
	res := dp[0]
	queue = append(queue, dp[0])

	for i := 1; i < k ; i++{
		dp[i] = max(nums[i],nums[i] + queue[0])
		res = max(res,dp[i])

		for len(queue) != 0 && queue[len(queue)-1] < dp[i]{

				queue = queue[0:len(queue)-1]

		}
		queue = append(queue, dp[i])
	}

	for i:= k ; i < len(nums) ; i++{
		dp[i] = max(nums[i],nums[i] + queue[0])
		res = max(res,dp[i])

		if dp[i-k] == queue[0]{
			queue = queue[1:]
		}
		for len(queue) != 0 && queue[len(queue)-1] < dp[i]{
			if len(queue) == 1{
				queue = []int{}
			}else{
				queue = queue[0:len(queue)-1]
			}
		}
		queue = append(queue, dp[i])
	}
	return res

}
func max(a,b int)int{
	if a > b{
		return a
	}else{
		return b
	}
}

func main() {

	a := []int{10,2,-10,5,20}
	fmt.Println(constrainedSubsetSum(a,2))
	
}
