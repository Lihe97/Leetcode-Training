package main

import "fmt"

func t1(nums [][]int)  {

	for i := 0 ; i < len(nums)/2 ; i ++{
		nums[i],nums[len(nums)-i-1] = nums[len(nums)-i-1],nums[i]
	}
	for i := 0 ; i < len(nums) ; i ++{
		for j := 0 ;j < len(nums[i])/2 ; j ++{
			nums[i][j],nums[i][len(nums[i])-j-1] = nums[i][len(nums[i])-j-1],nums[i][j]
		}
	}
	for i := 0 ; i < len(nums); i ++{
		for j := 0 ; j  < len(nums) ; j ++{
			fmt.Print(nums[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}


}

func main(){

	var n int
	fmt.Scan(&n)
	nums := make([][]int,n)
	for i := 0 ; i < n ; i ++{
		temp := make([]int,n)
		nums[i] = temp
	}
	for i := 0 ; i < n ; i ++{
		for j := 0 ; j < n ; j ++{
			fmt.Scan(&nums[i][j])
		}
	}

	t1(nums)

}
//3
//1 2 3
//4 5 6
//7 8 9
//4
//1 2 3 4
//5 6 7 8
//9 10 11 12
//13 14 15 16