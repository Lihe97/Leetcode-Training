package main

import "fmt"

func main() {

	a := 3

	b := &a

	c := *b

	fmt.Println(b,c)

}
