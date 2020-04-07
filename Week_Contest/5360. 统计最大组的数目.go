package main

import "fmt"

func countLargestGroup(n int) int {

	mp := map[int][]int{}
	for i := 1; i <= n ; i ++{
		mp[cal(i)] = append(mp[cal(i)], i)
	}
	max := 0
	res := 0
	for _,a := range mp{
		if len(a) > max{
			max = len(a)
			res = 1
		}else if len(a) == max{
			res ++
		}

	}
	return res
}
func cal(x int)int{
	res := 0
	for x != 0{
		res += x % 10
		x = x / 10
	}
	return res
}
func main() {
	fmt.Println(countLargestGroup(28))
	
}
