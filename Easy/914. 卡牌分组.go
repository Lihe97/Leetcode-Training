package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {

	mp := map[int]int{}
	for i := 0 ; i < len(deck) ; i ++{
		mp[deck[i]]++
	}
	g := 0
	for _,value := range mp{
		if g==0{
			g = value
		}
		g = gcd(g,value)
	}
	return g>1

}
func gcd(a,b int)int{

	if a % b != 0{
		a , b = b , a%b
	}
	return b
}
func main() {
	a := []int{1,1,1,2,2,2,3,3}
	fmt.Println(hasGroupsSizeX(a))
	
}
