package main

func removeDuplicates(nums []int) int {

	k := 0

	for i := 1 ; i < len(nums) ; i ++{
		if nums[i] != nums[i-1]{
			k ++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

func main() {
	
}
