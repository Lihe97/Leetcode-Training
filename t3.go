package main

import "fmt"

func t3(nums []int,a int)  {

	sum := 0

	for i := 0 ; i < len(nums) ; i ++{
		sum += nums[i]
	}
	if sum < 0 {
		sum = -sum
	}
	//fmt.Println("???",sum/a)

	if sum % a == 0{
		fmt.Println(sum/a)
	}else{
		fmt.Println(sum/a + 1)
	}

}

func main() {

	var n int
	var a int

	fmt.Scan(&n)
	fmt.Scan(&a)


	nums := make([]int,n)

	for i := 0 ; i < n ; i ++{
		fmt.Scan(&nums[i])
	}
	t3(nums,a)
}