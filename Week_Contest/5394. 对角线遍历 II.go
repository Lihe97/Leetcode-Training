package main

import "fmt"

type point struct{
	x int
	y int
}

func findDiagonalOrder(nums [][]int) []int {

	flag := [][]bool{}
	for i := 0 ; i < len(nums); i ++{
		temp := make([]bool,len(nums[i]))
		flag = append(flag, temp)
	}
	res := []int{}
	stack := []point{}
	stack = append(stack, point{0,0})
	flag[0][0] = true
	for len(stack) != 0{
		temp := stack[0]
		stack = stack[1:]
		res = append(res, nums[temp.x][temp.y])
		if temp.x + 1 < len(nums) && temp.y < len(nums[temp.x+1]) && !flag[temp.x+1][temp.y] {
			stack = append(stack, point{temp.x+1,temp.y})
			flag[temp.x+1][temp.y] = true
		}
		if temp.y + 1 < len(nums[temp.x]) && !flag[temp.x][temp.y+1]{
			stack = append(stack, point{temp.x,temp.y+1})
			flag[temp.x][temp.y+1] = true
		}
	}
	return res
}

func main() {

	a := [][]int{{1,2,3,4,5},{6,7},{8},{9,10,11},{12,13,14,15,16}}

	fmt.Println(findDiagonalOrder(a))
	
}
