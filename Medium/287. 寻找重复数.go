package main
func findDuplicate(nums []int) int {

	left , right := 1,len(nums)-1

	for left < right{
		mid := (left + right) / 2
		count := 0
		for i := 0 ; i < len(nums); i ++{
			if nums[i] <= mid {
				count ++
			}
		}
		if count <= mid {
			left = mid + 1
		}else{
			right = mid
		}
	}
	return left

}

func main() {
	
}
