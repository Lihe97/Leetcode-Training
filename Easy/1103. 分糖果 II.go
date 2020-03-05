package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {

	res := make([]int,num_people)
	i := 0
	a := 1
	for candies >0{
		i = (i+num_people)%num_people
		if candies > a{
			res[i] += a
			candies = candies-a
		}else{
			res[i] += candies
			break
		}
		a++
		i++
	}
	return res

}
func main() {
	fmt.Println(distributeCandies(10,3))
	
}
