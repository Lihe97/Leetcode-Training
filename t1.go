package main

import "fmt"

func containsPattern(arr []int, m int, k int) bool {

	left := 0
	right := m*k - 1

	for right < len(arr) {
		flag := true
		for j := 0  ; j  < m ; j ++{
			for i := left + m + j ; i <= right ; i+= m{
				if arr[i] != arr[i-m]{
					flag = false
					break
				}
			}
			if !flag{
				break
			}
		}
		if flag{
			return true
		}else{
			right ++
			left ++
		}
	}
	return false
}


func main() {

	fmt.Println(containsPattern([]int{1,2,1,2,1,3},2,3))


}