package main

import "fmt"

//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//
//	res := []int{}
//	if len(nums1) == 0{
//		return res
//	}
//	mp := map[int]int{}
//	temp := []int{}
//
//	temp = append(temp, nums2[0])
//	for i := 1 ; i < len(nums2) ; i ++{
//		if temp[0] >= nums2[i]{
//			temp = append(temp, nums2[i])
//		}else{
//
//			for  len(temp) != 0 && temp[0] < nums2[i]{
//				mp[temp[0]] = nums2[i]
//				if len(temp) == 1{
//					temp = []int{}
//				}else {
//					temp = temp[1 : len(temp)]
//				}
//			}
//			temp = append(temp, nums2[i])
//		}
//
//
//	}
//	if len(temp) != 0{
//		for _,v := range temp{
//			mp[v] = -1
//		}
//	}
//
//	for _,x := range nums1{
//		res = append(res, mp[x])
//	}
//
//	return res
//}
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	temp := []int{}
	mp := map[int]int{}
	res := []int{}

	for i := 0 ; i < len(nums2); i ++{

		for len(temp) != 0 && nums2[i] > temp[len(temp)-1]{
			mp[temp[len(temp)-1]] = nums2[i]
			temp = temp[0:len(temp)-1]
		}
		temp = append(temp, nums2[i])

	}
	if len(temp) != 0{
		for _,v := range temp{
			mp[v] = -1
		}
	}
	for i := 0 ; i  < len(nums1) ; i ++{
		res = append(res, mp[nums1[i]])
	}

	return res
}

func main() {


	nums1 := []int{4,1,2}
	nums2 := []int{1,3,4,2}
	fmt.Println(nextGreaterElement(nums1,nums2))
	
}
