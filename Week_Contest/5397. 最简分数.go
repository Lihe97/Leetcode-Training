package main

import (
	"fmt"
	"strconv"
)

func simplifiedFractions(n int) []string {

	res := []string{}

	if n == 1{
		return res
	}
	for i := 2 ; i <= n ; i++{

		//p := strconv.Itoa(1) + "/" + strconv.Itoa(i)
		//
		//res = append(res,p)
		//fmt.Println(res)

		for j := 1; j <= i-1; j ++{

			temp := false
			for k := 2 ; k <= j && temp == false; k++{
				if i % k == 0 && j % k == 0{
					temp = true
					break
				}
			}
			if !temp{
				q := strconv.Itoa(j) + "/" + strconv.Itoa(i)
				res = append(res,q)
			}

		}
	}

	return res

}


func main() {


	fmt.Println(simplifiedFractions(3))
	
}
