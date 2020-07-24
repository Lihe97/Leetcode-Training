package main

import (
	"fmt"
)

func t(s1 string,s2 string)  {
	if len(s1) < len(s2){
		fmt.Println(0)
		return
	}
	min := 1<<32

	for i := 0 ; i < len(s1) ; i ++{

			if s1[i] == s2[0]{
				a := i + 1
				b := 1
				for a < len(s1) && b < len(s2) {
					if s1[a] == s2[b]{
						b ++
					}
					a ++
				}
				//fmt.Println(a,b)
				if b == len(s2){
					if a - i < min{
						min = a - i
					}
				}
			}
	}
	if min == 1 <<32{
		fmt.Println(0)
		return
	}
	fmt.Println(min)

}




func main() {

	var s1 string
	var s2 string
	fmt.Scan(&s1)
	fmt.Scan(&s2)

	t(s1,s2)


}


