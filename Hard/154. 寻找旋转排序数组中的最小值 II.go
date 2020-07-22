package main
func findMin(nums []int) int {
	left := 0
	right := len(nums)-1

	for left < right{
		mid := (left + right) /2

		if nums[right] > nums[mid]{
			right = mid
		}else if nums[right] < nums[mid]{
			left = mid + 1
		}else{
			right --
		}

	}
	return nums[left]

}
func main() {
	
}
