package main

import (
	"fmt"
)

func t1(nums1 []int,nums2 []int,a int){
	min := 1<<32

	for i := 0 ; i < len(nums1) ; i ++{
		for j := 0 ; j < len(nums2) ; j++{
			if (nums1[i] + nums2[j]) % a < min{
				min = (nums1[i] + nums2[j]) % a
			}
		}
	}
	fmt.Println(min)
}

func main(){


	var m int
	var n int
	var a int


	fmt.Scan(&m)
	fmt.Scan(&n)
	fmt.Scan(&a)

	nums1 := make([]int,m)
	nums2 := make([]int,n)

	for i := 0 ; i < m ; i ++{
		fmt.Scan(&nums1[i])
	}
	for i := 0 ; i < n ; i ++{
		fmt.Scan(&nums2[i])
	}
	t1(nums1,nums2,a)


}
