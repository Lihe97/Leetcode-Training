package main

import "fmt"
func uniqueOccurrences(arr []int) bool {

	mp := map[int]int{}
	for i := 0 ; i < len(arr) ; i ++{
		mp[arr[i]] ++
	}
	times := make(map[int]int)
	for _, v := range mp {
		times[v]++
		if times[v] > 1 {
			return false
		}
	}
	return true

}

func main() {
	a := []int{1,2}
	fmt.Println(uniqueOccurrences(a))
	
}
