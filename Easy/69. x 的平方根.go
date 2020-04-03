package main

import "fmt"

//func mySqrt(x int) int {
//
//	left := 1
//	right := x/2
//	if x == 1{
//		return 1
//	}
//
//	for left <= right{
//		mid := (left + right)/2
//		if mid * mid == x {
//			return mid
//		}else  if mid * mid < x{
//			left  = mid +1
//		}else{
//			right = mid -1
//		}
//	}
//	return left-1
//
//}

func mySqrt(x int) int {

	if x == 1{
		return 1
	}
	m := x / 2
	for  m * m > x {
		m = (m + x/m)/2
	}

	return int(m)
}

func main() {

	fmt.Println(mySqrt(0))
}
