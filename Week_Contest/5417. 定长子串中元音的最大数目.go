package main

import "fmt"

func maxVowels(s string, k int) int {

	stack := []bool{}
	temp := []byte(s)

	res := 0
	for i := 0 ; i < k ; i ++{
		if temp[i] == 'a' || temp[i] == 'e' || temp[i] == 'i' || temp[i] == 'o' || temp[i] == 'u'{
			res ++
			stack = append(stack,true)
		}else{
			stack = append(stack,false)
		}
	}
	max := res

	right := k

	for right < len(temp){
		if stack[0] == true{
			res --
		}
		stack = stack[1:]
		if temp[right] == 'a' || temp[right] == 'e' || temp[right] == 'i' || temp[right] == 'o' || temp[right] == 'u'{
			res ++
			stack = append(stack,true)
		}else{
			stack = append(stack,false)
		}
		if res > max{
			max = res
		}
		right++
	}
	return max
}

func main() {

	fmt.Println(maxVowels("ibpbhixfiouhdljnjfflpapptrxgcomvnb",33))
	
}
