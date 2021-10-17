package main

import "fmt"

func getArr(arr []int) [][]int{
	res := [][]int{}

	end := 1 << len(arr)
	mark := 0
	for mark = 0 ; mark < end ; mark++{
		temp := []int{}
		for i := 0 ; i < len(arr) ; i ++{
			if ((1<<i)&mark)!=0 {

				temp = append(temp,arr[i])
			}
		}
		res = append(res,temp)
	}
	fmt.Println("?",res)
	return res
}
func maxAlternatingSum(nums []int) int64 {

	arr := getArr(nums)
	var max int64
	for i := 1 ; i < len(arr) ; i ++{

		var temp int64
		for j := 0 ; j < len(arr[i]) ; j ++{

			if j % 2 == 0{
				temp += int64(arr[i][j])
			}else{
				temp -= int64(arr[i][j])
			}
		}
		if temp > max{
			max = temp
		}
	}
	return max
}
func main() {
	fmt.Println(maxAlternatingSum([]int{374,126,84,237,195,139,328,353,286,113,351,167,394,398,29,118,17,162,206,138,34,109,291,368,162,109,336,256,203,330,235,74,136,72,127,382,288,276,135,383,300,220,299,205,186,113,71,261,253,47,387,25,57,79,322,82,349,217,306,33,198,196,306,240,271,129,284,6,349,370,59,350,275,385,137,394,329,175,58,151,223,81,233,2,370,369,135,257,391,92,260,55,321,153,328,260,312,102,79,192}))
}
