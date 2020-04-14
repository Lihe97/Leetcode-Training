package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {

	left := 0
	right := len(letters)-1

	for left < right{
		mid := (left + right) / 2
		if letters[mid] > target{
			right = mid
		}else{
			left = mid + 1
		}
	}
	if letters[left] > target{
		return letters[left]
	}else{
		return letters[0]
	}


}

func main() {
	a := []byte{'c','f','j'}
	fmt.Println(nextGreatestLetter(a,'d'))

	
}
