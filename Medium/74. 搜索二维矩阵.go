package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) <  1{
		return false
	}

	row := getRow(matrix,target)
	//fmt.Println(row)
	return find(matrix[row],target)
}
func  find( nums []int, target int) bool {

	left := 0
	right := len(nums) -1

	for left <= right{
		mid := (left + right ) / 2
		if nums[mid] == target{
			return true
		}
		if nums[mid] < target{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return false

}
func getRow(matrix [][]int, target int) int {

	top := 0
	bottom := len(matrix) - 1
	col := len(matrix[0]) - 1

	for top < bottom{

		mid := (top + bottom) / 2

		if matrix[mid][col] < target{
			top = mid + 1
		}else{
			bottom = mid
		}
	}
	return top
}

func main() {

	a := [][]int{{1}}
	fmt.Println(searchMatrix(a,1))

}
