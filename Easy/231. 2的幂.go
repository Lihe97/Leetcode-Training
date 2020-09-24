package main

import "fmt"

func isPowerOfTwo(n int) bool {

	sum := 0

	for n != 0 {
		sum += n & 1
		if sum > 1{
			return false
		}
		n = n >> 1
	}
	return sum == 1

}

func isPowerOfTwo(n int) bool {


	return n > 0 && n & (n-1) == 0

}
func main() {
	fmt.Println(isPowerOfTwo(32))
}
