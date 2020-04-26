package main

import "fmt"

type point struct{
	x int
	y int
}

func findDiagonalOrder(nums [][]int) []int {
	maxy := 0

	for i := 0 ; i < len(nums) ; i ++{
		if len(nums[i]) > maxy{
			maxy = len(nums[i])

		}
	}
	res := []int{}

	stack := []point{}
	for i := 0 ; i < len(nums)-1; i++{
		stack = append(stack, point{i,0})
	}
	for i := 0 ; i < len(nums[len(nums)-1]) ; i++{
		stack = append(stack, point{len(nums)-1,i})
	}
	for len(stack) != 0{
		temp := stack[0]
		stack = stack[1:]
		for temp.x >= 0 && temp.y < maxy {
			if temp.y < len(nums[temp.x]){
				res = append(res, nums[temp.x][temp.y])
			}
			temp.x --
			temp.y ++
		}
	}

	lastx := len(nums)-1
	lasty := len(nums[len(nums)-1])


	for i := lastx - 1; i >= 0 ; i--{

		for j := lastx - i +1 ; j < len(nums[i]) ; j ++{

			if (j - lasty) !=( i - lastx){
				res = append(res, nums[i][j])
			}
		}
	}



	return res

}

func main() {

	a := [][]int{{1,2,3,4,5},{6,7},{8},{9,10,11},{12,13,14,15,16}}

	fmt.Println(findDiagonalOrder(a))
	
}
