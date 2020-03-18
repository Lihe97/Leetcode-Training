package main

import "fmt"
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return max(rec1[0], rec2[0]) < min(rec1[2], rec2[2]) && max(rec1[1], rec2[1]) < min(rec1[3], rec2[3])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	a:= []int{2,17,6,20}
	b:= []int{3,8,6,20}
	fmt.Println(isRectangleOverlap(a,b))
	
}
