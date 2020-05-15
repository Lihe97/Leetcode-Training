package main

import "fmt"

//func subarraySum(nums []int, k int) int {
//
//	qz := make([]int,len(nums)+1)
//	res := 0
//	qz[0] = 0
//	for i := 0 ; i < len(qz) ; i ++{
//		qz[i+1] = nums[i] + qz[i]
//	}
//
//
//	for i := 0 ; i < len(qz) ; i++{
//		for j := i; j < len(qz) ; j++{
//			if qz[j+1] - qz[i] == k{
//				res ++
//			}
//		}
//	}
//	return res
//}

func subarraySum(nums []int, k int) int {
	mp := map[int]int{}
	mp[0] = 1

	res , sum := 0,0

	for i :=  0; i  < len(nums)  ; i ++{
		sum += nums[i]

		if 	a,ok := mp[sum - k];ok{
			res += a
		}
		mp[sum] ++
	}

	return res
}

func main() {

	a := []int{-1,-1,1}
	fmt.Println(subarraySum(a,0))
	
}
