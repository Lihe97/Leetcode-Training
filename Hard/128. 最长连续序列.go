package main

import "fmt"

func longestConsecutive(nums []int) int {

	ans := 0



	mp := map[int]bool{}

	for i := 0 ; i < len(nums) ; i ++{
		mp[nums[i]] = true
	}
	for a := range mp{

		if !mp[a-1]{
			cur := a
			temp := 1
			for mp[cur+1]{
				cur ++
				temp ++
			}
			if temp > ans {
				ans =temp
			}
		}
	}
	return ans

}

func main() {


	a := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(a))
}
