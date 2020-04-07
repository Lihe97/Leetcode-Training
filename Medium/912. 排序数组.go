package main

import "fmt"

//func sortArray(nums []int) []int {
//
//	sort.Ints(nums)
//	return nums
//}
func sortArray(nums []int) []int {

	quicksort(nums,0,len(nums)-1)
	return nums
}
func quicksort(arr []int, left, right int) []int {
	if left < right {
		p := ppartition(arr, left, right)
		quicksort(arr, left, p - 1)
		quicksort(arr, p + 1, right)
	}
	return arr
}
func ppartition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
	return index - 1
}



func main() {
	a := []int{49,38,65,97,76,13,27}
	fmt.Println(sortArray(a))
	
}
