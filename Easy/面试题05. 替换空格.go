package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.ReplaceAll(s," ","%20")

}
func main() {
	fmt.Println(replaceSpace("we are one"))
	
}
