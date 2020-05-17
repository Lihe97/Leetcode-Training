package main

import (
	"fmt"
	"sort"
	"strings"
)

func arrangeWords(text string) string {

	temp := strings.Split(text," ")

	temp[0]  = strings.ToLower(temp[0])

	sort.SliceStable(temp, func(i, j int) bool {
		return len(temp[i]) < len(temp[j])
	})

	temp[0] = strings.Title(temp[0])


	return strings.Join(temp," ")

}

func main() {

	fmt.Println(arrangeWords("abc ab c b"))
}
