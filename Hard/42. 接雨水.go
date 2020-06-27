package main

import "fmt"

//func trap(height []int) int {
//
//	stack := []int{}
//	res := 0
//	for i := 0 ; i < len(height) ; i ++{
//
//		for len(stack) != 0 && height[stack[len(stack) - 1]] < height[i]{
//			curIndex := stack[len(stack)-1]
//			stack = stack[:len(stack)-1]
//			for len(stack) != 0 && height[stack[len(stack)-1]] == height[curIndex]{
//				stack = stack[:len(stack)-1]
//			}
//
//			if len(stack) != 0 {
//
//				res += (min(height[i], height[stack[len(stack)-1]]) - height[curIndex]) * (i - stack[len(stack)-1] - 1)
//			}
//		}
//		stack = append(stack, i)
//
//	}
//	return res
//}
//func min(a,b int)int{
//	if a < b{
//		return a
//	}else{
//		return b
//	}
//}
func trap(height []int) int {
	stack := []int{}
	res := 0
	for i := 0 ; i < len(height) ; i ++{
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]]{
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for len(stack) != 0 && stack[len(stack)-1] == height[cur]{
				stack = stack[:len(stack)-1]
			}
			if len(stack) != 0{
				res +=  (min(height[i],height[stack[len(stack)-1]]) - height[cur]) * (i - stack[len(stack)-1] -1 )
			}
		}
		stack = append(stack,i)
	}
	return res
}
func main() {

	a := []int{}
	fmt.Println(trap(a))
	
}
