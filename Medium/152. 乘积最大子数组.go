package main

import "fmt"


func maxProduct(nums []int) int {

	dpmin := make([]int,len(nums))
	dpmax := make([]int,len(nums))
	dpmin[0] = nums[0]
	dpmax[0] = nums[0]
	res := nums[0]
	for i := 1 ; i < len(nums) ; i ++{
		if nums[i] < 0 {
			dpmax[i] = max(nums[i] * dpmin[i-1],nums[i])
			dpmin[i] = min(nums[i] * dpmax[i-1],nums[i])
		}else{
			dpmax[i] = max(nums[i] * dpmax[i-1],nums[i])
			dpmin[i] = min(nums[i] * dpmin[i-1],nums[i])
		}
		if dpmax[i] > res{
			res = dpmax[i]
		}
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
	a := []int{2,3,-2,4}
	fmt.Println(maxProduct(a))
	
}
