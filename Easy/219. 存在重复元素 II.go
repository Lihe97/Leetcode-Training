package main

import (
	"fmt"
)

//func containsNearbyDuplicate(nums []int, k int) bool {
//	mp := map[int][]int{}
//
//	for i := 0 ; i < len(nums) ; i++{
//		mp[nums[i]] = append(mp[nums[i]], i)
//	}
//	//fmt.Println(mp)
//	for _,b := range mp{
//		if len(b) > 1{
//			sort.Ints(b)
//			for i := 1;  i < len(b) ; i++{
//				if b[i] - b[i-1] <= k{
//					return true
//				}
//			}
//		}
//	}
//	return false
//}
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := map[int]int{}

	for i := 0 ; i < len(nums) ; i ++{

		if _,ok  := mp[nums[i]]; ok{
			return true
		}
		mp[nums[i]] = i

		if len(mp)  > k{
			delete(mp,nums[i-k])
		}
	}
	return false
}
func main() {

	a := []int{1,2,3,1}
	fmt.Println(containsNearbyDuplicate(a,3))
	
}
