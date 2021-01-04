package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {



	mp := map[string]bool{}

	for _,w := range wordList{
		mp[w] = true
	}
	if mp[endWord] == false{
		return 0
	}

	queue := []string{}
	queue = append(queue,beginWord)
	level := 1

	for len(queue) != 0{

		temp := queue
		queue = []string{}
		for i := 0 ; i < len(temp) ; i ++{
			word := temp[i]
			if word == endWord{
				return level
			}
			for a,b := range mp{
				if compare(word,a) && b{
					queue = append(queue,a)
					mp[a] = false
				}
			}
		}
		//fmt.Println(level,queue)
		level ++
	}

	return 0

}

func compare(a,b string) bool{
	sum := 0
	for i := 0 ; i < len(a)  ; i ++{
		if a[i] != b[i] {
			sum ++
		}
	}
	return sum == 1
}
func main() {


	fmt.Println(ladderLength("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
}
