package main

import "fmt"

func t2(nums []int){
	left := 0
	right := 0
	max := -1

	maxl := -1
	maxr := -1

	for right < len(nums) - 1 {
		if nums[right+1] > nums[right]{
			right ++
			if right - left + 1 > max{
				maxr = right
				maxl = left
				max = right - left + 1
			}
		}else{
			right ++
			left = right
		}
		//fmt.Println("??")
		//fmt.Println(right,maxl,maxr)
	}
	if maxr == -1{
		fmt.Print("N")
		return
	}
	temp := nums[maxl:maxr+1]
	for i := 0 ; i  < len(temp) ;  i ++{
		fmt.Print(temp[i])
		fmt.Print(" ")
	}
	//fmt.Print(nums[maxl:maxr+1])

}

func main() {

	var n int
	fmt.Scan(&n)
	nums := make([]int,n)
	for i := 0 ; i < n ; i ++{
		fmt.Scan(&nums[i])
	}
	t2(nums)
}