package main

import "fmt"

func waysToStep(n int) int {
	return fb(n) % 1000000007

}
func fb(n int ) int{
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	first := 4
	second := 2
	third := 1
	for i:= 3; i < n ; i++{
		first,second,third = (first + second + third)% 1000000007, first% 1000000007, second% 1000000007
	}
	return first
}
func main() {
	fmt.Println(waysToStep(76))
	
}
