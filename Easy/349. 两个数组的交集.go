package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]bool{}
	m2 := map[int]bool{}
	res := []int{}

	for i := 0 ; i <len(nums1);i++{
		m1[nums1[i]] = true
	}
	for i := 0 ; i <len(nums2);i++{
		m2[nums2[i]] = true
	}
	for x,_ := range m1{
		if m1[x] == m2[x]{
			res = append(res, x)
		}
	}
	return res
}


func main() {
	a1 := []int{4,9,5}
	a2 := []int{9,4,9,8,4}
	fmt.Println(intersection(a1,a2))

	
}
