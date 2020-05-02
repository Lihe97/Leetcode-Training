package main
func kidsWithCandies(candies []int, extraCandies int) []bool {

	max := 0
	res := make([]bool,len(candies))
	for i := 0 ; i < len(candies) ; i ++{
		if candies[i] > max{
			max = candies[i]
		}
	}
	for i := 0 ; i < len(candies) ; i++{
		if candies[i] + extraCandies >= max{
			res[i] = true
		}else{
			res[i] = false
		}
	}
	return res

}

func main() {
	
}
