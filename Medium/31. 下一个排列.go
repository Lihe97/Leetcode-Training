package main

func nextPermutation(nums []int)  {

	i := len(nums) - 2
	j := len(nums) - 1
	for i >= 0 && nums[i] >= nums[j]{
		i --
		j --
	}
	if i >= 0{
		k := len(nums) - 1
		for nums[k] <= nums[i]{
			k --
		}
		nums[k],nums[i] = nums[i],nums[k]
	}
	m,n := j,len(nums)-1

	for m < n {
		nums[m],nums[n] = nums[n],nums[m]
		m ++
		n --
	}


}
func main() {
	
}
