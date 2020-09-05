package main
func minOperations(nums []int) int {

	res := 0

	ji,ou := 0,0
	zero := 0

	for zero != len(nums){
		ji, ou = 0,0
		zero = 0
		for i := 0 ; i < len(nums) ; i ++{
			if nums[i] % 2 == 0{
				if nums[i] == 0{
					zero ++
				}
				ou ++
			}else{
				ji ++
			}

		}
		if zero == len(nums) {
			break
		}
		if ji == 0{
			res ++
			for i := 0 ; i < len(nums) ; i ++{
				nums[i] /= 2
			}
		}else{
			res += ji
			for i := 0 ; i < len(nums) ; i ++{
				if nums[i] % 2 == 1{
					nums[i] --
				}
			}
		}

	}


	return res
}
func main() {
	
}
