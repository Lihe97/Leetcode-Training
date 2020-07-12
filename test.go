package main

import (
	"fmt"
)

func patterRec(pattern string,str string) bool{

	temp := []string{}
	cur := 0
	for i := 0 ; i < len(str);  i ++{
		if str[i] == ' '{
			temp = append(temp,str[cur:i])
			cur = i + 1
		}
	}
	temp = append(temp,str[cur:len(str)])
	if len(pattern) != len(temp) {
		return false
	}
	mp := map[byte]string{}
	for i := 0 ; i < len(pattern) ; i ++{
		if _,ok := mp[pattern[i]];ok{
			if mp[pattern[i]] != temp[i]{
				return false
			}
		}else{
			mp[pattern[i]] = temp[i]
		}
	}
	return true
}



func main() {

	fmt.Println(10&(-10))


}


