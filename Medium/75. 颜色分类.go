package main

func sortColors(nums []int)  {

	p0,p1 := 0,0

	for i := 0 ; i  < len(nums)  ; i ++{
		if nums[i] == 1{
			nums[i],nums[p1] = nums[p1],nums[i]
			p1 ++

		}else if nums[i] == 0 {
			nums[p0],nums[i] = nums[i],nums[p0]
			if p0 < p1{
				nums[p1],nums[i] = nums[i],nums[p1]
			}
			p0 ++
			p1 ++
		}
	}
}

func main() {
	
}
