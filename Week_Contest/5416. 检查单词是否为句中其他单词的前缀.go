package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {

	temp := strings.Split(sentence," ")

	for i := 0 ; i < len(temp) ; i++{
		if strings.HasPrefix(string(temp[i]),searchWord){
			return i
		}
	}
	return -1

}

func main() {

}
