package main

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {


	arr := strings.Split(s," ")

	temp := []int{}

	for i := 0 ; i  < len(arr) ; i ++{
		if arr[i][0] >= '0' && arr[i][0] <= '9'{
			tmp,_ := strconv.Atoi(arr[i])
			temp = append(temp,tmp)
		}
	}
	for i := 1 ; i  < len(temp) ; i ++{
		if temp[i] <= temp[i-1]{
			return false
		}
	}
	return true

}
