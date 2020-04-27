package main
func search(nums []int, target int) int {

	low,high := 0,len(nums)-1

	for low <= high{
		mid := (low + high)/2
		if nums[mid] == target{
			return mid
		}
		if nums[mid] >= nums[low]{
			if nums[mid] > target && target >= nums[low]{
				high = mid - 1
			}else{
				low = mid + 1
			}
		}else{
			if nums[mid] < target && target <= nums[high]{
				low = mid +1
			}else{
				high = mid - 1
			}

		}
	}
	return -1


}

func main() {
	
}
