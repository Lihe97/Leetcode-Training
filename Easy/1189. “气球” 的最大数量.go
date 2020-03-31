package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	mp := map[byte]int{}
	mp['a'] = 0
	mp['b'] = 0
	mp['l'] = 0
	mp['o'] = 0
	mp['n'] = 0

	for i := 0 ; i < len(text) ; i ++{
		if text[i] == 'b' || text[i] =='a' || text[i] == 'n' || text[i] == 'l' || text[i] =='o'{
			mp[text[i]]++
		}

	}

	min := 10<<32
	for a,v := range mp{
		if a == 'b' || a =='a' || a == 'n'{
			if v < min{
				min = v
			}
		}else if a == 'l' || a =='o'{
			if v/2 < min{
				min = v/2
			}
		}

	}
	return min
}
func main() {
	fmt.Println(maxNumberOfBalloons("nlaebloko"))

	
}
