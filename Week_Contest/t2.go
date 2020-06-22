package main

import (
	"fmt"
	"strconv"
)

func getFolderNames(names []string) []string {
	res := []string{}

	mp := map[string]int{}

	for i := 0 ; i < len(names) ; i ++{

		temp := names[i]

		if mp[temp] == 0{
			res = append(res,temp)
			mp[temp] ++
		}else{
			t := temp[:len(temp)-3]
			if temp[len(temp)-1] != ')'  && mp[temp] != 0 {
				mp[temp]++
				temp = temp + "(" + strconv.Itoa(mp[temp]) + ")"
				res = append(res,temp)
				mp[temp] ++
			}else if temp[len(temp)-1] == ')' && mp[t] != 0{
				temp = temp + "(" + strconv.Itoa(mp[t]) + ")"
				res = append(res,temp)
				//mp[temp] ++
				mp[t] ++/
			}else if  temp[len(temp)-1] == ')' && mp[temp] == 0{
				temp = temp + "(" + strconv.Itoa(mp[temp]) + ")"
				res = append(res,temp)
				mp[temp] ++
			}


		}
	}
	fmt.Println(mp)
	return res
}
func main() {

	a := []string{"kaido","kaido(1)","kaido","kaido(1)"}
	fmt.Println(getFolderNames(a))
	
}
