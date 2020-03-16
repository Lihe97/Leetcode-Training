package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {

	temp := []byte(S)
	var res string
	if len(S) == 1{
		return S
	}
	count := 1
	for i := 1 ;i < len(temp) ; i++{
		if temp[i] == temp[i-1]{
			count ++
			if i == len(temp)-1{
				res += string(temp[i])
				res += strconv.Itoa(count)
			}
		}else{
			res += string(temp[i-1])
			res += strconv.Itoa(count)
			count = 1
			if i == len(temp)-1{
				res += string(temp[i])
				res += strconv.Itoa(count)
			}
		}
	}
	if len(res)>=len(S){
		return S
	}else {
		return res
	}
}
func main() {
	fmt.Println(compressString("aa"))
}
