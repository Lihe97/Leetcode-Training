package main

func findKthPositive(arr []int, k int) int {

	temp := make([]bool,100001)
	for i:= 0 ; i < len(arr) ; i ++{
		temp[arr[i]] = true
	}
	cur := 0
	for i := 1 ; i < len(temp) ; i++{
		if temp[i] == false{
			cur ++
			if cur == k{
				return i
			}
		}
	}
	return 0
}

func main() {




}