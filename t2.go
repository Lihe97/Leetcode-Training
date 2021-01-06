package main

import "fmt"

func countPairs(d []int) int {

	res := 0
	mod := 1e9 + 7

	mp := map[int]int{}

	for i := 0 ; i < len(d) ; i ++{

		for j := 0 ; j <= 21 ; j ++{
			temp := (1 << j) - d[i]
			if _,ok := mp[temp];ok{
				res += mp[temp]
			}
		}
		mp[d[i]] ++
	}
	return res % int(mod)
}
func main() {

	fmt.Println(countPairs([]int{1,3,5,7,9}))


}