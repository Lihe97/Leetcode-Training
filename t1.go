package main

func smallestEqual(nums []int) int {
	res := -1
	for i := 0 ; i < len(nums) ; i ++{
		if i % 10 == nums[i]{
			res = i
		}
		if res != -1{
			break
		}
	}
	return res

}


func main(){



}
