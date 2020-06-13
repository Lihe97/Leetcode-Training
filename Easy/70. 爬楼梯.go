package main

import "fmt"

//func climbStairs(n int) int {
//	return climb(n)
//}
//func climb(n int) int {
//	if n == 1 {
//		return 1
//	}230
//	one, two := 1, 2
//	for i := 3; i <= n; i++ {
//		one, two = two, one + two
//	}
//	return two
//}

func climbStairs(n int) int {

	temp := make([]int,n + 1)
	temp[0] = 1
	temp[1] = 1


	for i := 2 ; i <= n ; i ++{
		temp[i] = temp[ i - 1 ] + temp[ i-2 ]

	}
	//fmt.Println(temp)
	return temp[n]

}

func main() {

	fmt.Println(climbStairs(3))
	
}
