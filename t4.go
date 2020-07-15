package main

import "fmt"

func main() {
	//var a byte

	b := make([]int,2)
	for i := 0 ; i < len(b) ; i ++{
		fmt.Scanf("%d",&b[i])
	}
	fmt.Println(b)

}
