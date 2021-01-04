package main

import "math"

func countPairs(deliciousness []int) int {

	res := 0

	mp := map[int]int{}


	for i := 0 ; i < len(deliciousness) ; i ++{


		for j := 0 ; j <= 22 ; j ++{
			target := int(math.Pow(2,float64(j)))

			if target - deliciousness[i] < 0{
				continue
			}
			if _,ok := mp[target - deliciousness[i]];ok{
				res += mp[target - deliciousness[i]]
			}
		}
		mp[deliciousness[i]] ++
	}

	return res % (int(math.Pow10(9)) + 7)
}


func main() {
	
}
