package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {

	queue := []int{}
	res := []int{}
	queue = append(queue, nums[0])

	for i := 1 ; i < k ; i ++{

		for len(queue) != 0 && queue[len(queue)-1] < nums[i]{
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res = append(res, queue[0])
	for i := k ; i < len(nums); i ++{
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		for len(queue) != 0 && queue[len(queue)-1] < nums[i]{
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		res = append(res,queue[0])
	}
	return res

}

func main() {

	a := []int{7,2,4}
	fmt.Println(maxSlidingWindow(a,2))
}
