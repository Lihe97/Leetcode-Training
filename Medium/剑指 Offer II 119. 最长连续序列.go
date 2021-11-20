package main


func longestConsecutive(nums []int) int {

	if len(nums) == 0{
		return 0
	}

	mp := map[int]bool{}
	for i := 0 ; i < len(nums) ; i ++{
		mp[nums[i]] = true
	}
	res := 1
	for k,_ := range mp{
		if !mp[k-1]{
			cur := 1
			temp := k + 1
			for mp[temp]{
				temp ++
				cur ++
			}
			if cur > res{
				res = cur
			}
		}
	}
	return res
}