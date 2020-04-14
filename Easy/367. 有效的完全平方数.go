package main

import "fmt"

func isPerfectSquare(num int) bool {

	if num <= 1{
		return true
	}
	left, right := 1, num/2

	for left <= right{
		mid := (left + right)/2

		if mid * mid == num{
			return true
		}
		if mid * mid > num{
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return false
}

func main() {

	fmt.Println(isPerfectSquare(9))
	
}
