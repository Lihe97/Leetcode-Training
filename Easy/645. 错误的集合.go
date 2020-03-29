package main


func findErrorNums(nums []int) []int {

	res := []int{}
	x,y := 0,0
	mp  := map[int]int{}
	for i := 0 ; i < len(nums) ; i++{
		mp[i+1] = 0
	}
	for i := 0 ; i < len(nums) ; i++{
		mp[nums[i]] ++
	}
	for a,b := range mp{
		if b == 2{
			x = a
		}
		if b == 0{
			y = a
		}
	}
	res = append(res, x)
	res = append(res, y)

	return res
}
func main() {
	
}
