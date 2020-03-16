package main

import (
	"fmt"
)

func countPrimes(n int) int {
	isPrime := make([]bool,n)
	for i := 2;i * i <= n;i++ {
		if !isPrime[i]{
			for j := i * i;j < n;j += i {
				isPrime[j]= true
			}
		}
	}
	count := 0
	for k:= 2;k < n;k++ {
		if !isPrime[k] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(2))
	
}
