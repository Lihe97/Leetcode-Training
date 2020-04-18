package main

import "fmt"

//func maxArea(height []int) int {
//
//	stack := []int{}
//	max := 0
//
//	for i := 0 ; i < len(height) ; i ++{
//
//
//		for j := 0 ; j < len(stack) ; j ++{
//
//			if stack[j] < height[i]{
//				temp := stack[j] * (i-j)
//				if temp > max{
//					max = temp
//				}
//			}else {
//				temp := height[i] * (i-j)
//				if temp > max{
//					max = temp
//				}
//			}
//		}
//		stack = append(stack, height[i])
//	}
//	return max
//}
func maxArea(height []int) int {

	max := 0
	left := 0
	right := len(height) - 1

	for left < right {
		if height[left] <= height[right] {
			temp := (right - left) * height[left]
			if temp > max {
				max = temp
			}
			left++
		} else {
			temp := (right - left) * height[right]
			if temp > max {
				max = temp
			}
			right--
		}
	}
	return max
}

func main() {

	a := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(a))
	
}
