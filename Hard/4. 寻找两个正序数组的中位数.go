package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	mid := (len1 + len2) / 2
	cur1 , cur2 := 0,0
	cnt := 0
	temp := []int{}
	for cnt <= mid && cur1 < len1 && cur2 < len2{
		if nums1[cur1] <= nums2[cur2]{
			temp = append(temp,nums1[cur1])
			cur1 ++
		} else{
			temp = append(temp,nums2[cur2])
			cur2 ++
		}
		cnt ++
	}
	if cur1 == len1{
		for cnt <= mid{
			temp = append(temp,nums2[cur2])
			cur2 ++
			cnt++
		}
	}else if cur2 == len2{
		for cnt <= mid{
			temp = append(temp,nums1[cur1])
			cur1 ++
			cnt++
		}
	}
	//fmt.Println(temp)
	if (len1 + len2) % 2 == 0 {
		return (float64(temp[len(temp) - 1] + temp[len(temp) - 2]))/2
	}else{
		return float64(temp[len(temp)-1])
	}
}

func main() {

	a := []int{1,2}
	b := []int{3,4}
	fmt.Println(findMedianSortedArrays(a,b))
	
}
