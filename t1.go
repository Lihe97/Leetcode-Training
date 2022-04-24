package main

import "fmt"

func solution(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(process(nums[1:]), process(nums[:len(nums)-1]))
}
func process(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	a := []int{2, 3, 2}
	b := []int{1, 2, 3, 1}
	c := []int{2}
	d := []int{2, 3}
	e := []int{4, 2, 3, 6}
	fmt.Println(solution(a))
	fmt.Println(solution(b))
	fmt.Println(solution(c))
	fmt.Println(solution(d))
	fmt.Println(solution(e))

}
