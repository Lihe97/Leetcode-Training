package main

import "fmt"

func minInsertions(s string) int {

	t := 0
	res := 0

	for i := 0 ; i < len(s) ; i ++ {
		if s[i] == '('{
			if t >= 0{
				if t % 2 == 0{
					t +=2
				}else{
					res += 1
					t += 1
				}
			}else{
				temp := -t
				if temp % 2 == 0{
					res += temp/2
				}else{
					res += temp/2 + 2
				}
				t = 0
				t += 2
			}
		} else if s[i] == ')'{
			t -= 1
		}

	}

	if t > 0{
		t -= 1
		res += 1
	}

	if t > 0{
		res += t
	}else if t == 0{
		return res
	}else{
		temp := -t
		if temp % 2 == 0{
			res += temp/2
		}else{
			res += temp/2 + 2
		}
	}
	return res
}

func main() {

	//fmt.Print(minInsertions("))))("))
	fmt.Print(minInsertions("(((()(()((())))(((()())))()())))(((()(()()((()()))"))



}