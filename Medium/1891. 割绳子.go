package main

func maxxLength(ribbons []int, k int) int {

	left := 0
	right := 100000

	for left < right {
		mid := ( left + right + 1 ) / 2
		cnt := 0
		for i :=  0 ; i < len(ribbons) ; i ++{
			cnt += ribbons[i] / mid
		}
		if cnt < k {
			right = mid  -1
		}else{
			left = mid
		}
	}
	return left


}