package main

import (
	"fmt"
	"sort"
)
func tt2(n int,nums [][]int,c0,d0 int) int{
	res := 0
	sort.Slice(nums, func(i, j int) bool {
		if float64(nums[i][3])/float64(nums[i][2]) == float64(nums[j][3])/float64(nums[j][2]){
			return nums[i][2] < nums[j][2]
		}
		return float64(nums[i][3])/float64(nums[i][2]) > float64(nums[j][3])/float64(nums[j][2])
	})
	for i := 0 ; i < len(nums) ; i ++{
		if n <= 0 {
			break
		}
		temp1 := nums[i][0]/nums[i][1]
		temp2 := n/nums[i][2]
		t := min(temp1,temp2)
		if t * nums[i][3] < n/c0 * d0{
			break
		}
		res += t*nums[i][3]
		n -= t*nums[i][2]
	}
	if n != 0{
		res = res +  n/c0 * d0
	}
	return res
}
func min(a,b int) int  {
	if a < b {
		return a
	}else{
		return b
	}
}
func main() {

	var n int
	var m int
	var c0 int
	var d0 int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&c0)
	fmt.Scan(&d0)
	nums := [][]int{}


	for i := 0 ; i < m ; i ++{
		temp := make([]int,4)
		for j := 0 ; j < 4; j ++{
			fmt.Scan(&temp[j])
		}
		nums = append(nums,temp)
	}
	fmt.Println(tt2(n,nums,c0,d0))

}
