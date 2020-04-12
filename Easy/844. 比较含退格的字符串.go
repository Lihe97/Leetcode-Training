package main

import "fmt"

func backspaceCompare(S string, T string) bool {

	var data1 []byte = []byte(S)
	stack1 := []string{}
	for i := 0 ; i < len(data1) ; i ++{
		if string(data1[i]) == "#" && len(stack1) != 0{
			stack1 = stack1[:len(stack1)-1]
		}else if string(data1[i]) != "#"{
			stack1 = append(stack1, string(data1[i]))
		}
	}
	var data2 []byte = []byte(T)
	stack2 := []string{}
	for i := 0 ; i < len(data2) ; i ++{
		if string(data2[i]) == "#" && len(stack2) != 0{
			stack2 = stack2[:len(stack2)-1]
		}else if string(data2[i]) != "#"{
			stack2 = append(stack2, string(data2[i]))
		}
	}
	if len(stack1) != len(stack2){
		return false
	}
	for i := 0 ; i < len(stack1); i ++{
		if stack1[i] != stack2[i]{
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(backspaceCompare("ab#c","##ad#c"))
	
}
