package main

import "fmt"
func countCharacters(words []string, chars string) int {


	t := make([]int,26)
	res := 0
	for _ , x := range chars{
		t[x - 'a']++
	}

	for i := 0 ; i < len(words) ; i++{
		a := make([]int,26)
		copy(a,t)
		flag := true
		for j := range words[i]{
			if a[words[i][j]- 'a']<=0{
				flag = false
				break
			}else{
				a[words[i][j] - 'a'] --
			}
		}
		if flag{
			res += len(words[i])
		}
	}
	return res

}

func main() {
	a:=[]string{"cat","bt","hat"}
	b := "atach"
	fmt.Println(countCharacters(a,b))

}
