package main

import "fmt"

type node struct {
	x int

}
func main() {

	cache := map[int]int{}
	cache[1] = 1
	fmt.Println(cache)
	delete(cache,1)
	fmt.Println(cache)

}
