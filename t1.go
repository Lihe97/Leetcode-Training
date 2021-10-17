package main

import "fmt"

func canBeIncreasing(nums []int) bool {

	cnt := 0
	flag := false
	for i := 0 ; i <= len(nums) - 1  ;i ++{
		temp := make([]int,len(nums[:i]))
		copy(temp,nums[:i])

		temp = append(temp,nums[i+1:]...)
		flag = isDz(temp,len(nums)-1)
		//fmt.Println(flag,cnt)
		if flag {
			cnt ++
		}
	}
	return cnt >= 1

}
func isDz(a []int,n int)bool{
	if n == 0 || n == 1{
		return true
	}
	if n == 2 {
		return a[n-1] > a[n-2]
	}
	return isDz(a,n-1) && (a[n-1] > a[n-2])
}

func main(){

	a := []int{2,1,3}
	fmt.Println(canBeIncreasing(a))

}
