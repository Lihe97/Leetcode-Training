package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {

	m1 := map[int]int{}
	m2 := map[int]int{}
	res := []int{}

	for i := 0 ; i <len(nums1);i++{
		m1[nums1[i]] ++
	}
	for i := 0 ; i <len(nums2);i++{
		m2[nums2[i]] ++
	}
	for x,_ := range m1{
		if  m2[x]>0{
			for i := 0 ; i <min(m1[x],m2[x]);i++{
				res = append(res, x)
			}
		}
	}
	return res
}


func min(a,b int)int{
	if a<= b{
		return a
	}else {
		return b
	}
}
func main() {
	a1 := []int{4,9,5}
	a2 := []int{9,4,9,8,4}
	fmt.Println(intersect(a1,a2))
	
}
