package main

func finalPrices(prices []int) []int {

	res := make([]int,len(prices))
	for i := 0 ;  i < len(prices) ; i++{
		res[i] = prices[i]
		for j := i + 1; j < len(prices) ;j++{
			if prices[j] <= prices[i]{
				res[i] = prices[i] - prices[j]
				break
			}
		}
	}
	//res[len(prices)-1] = prices[len(prices)-1]
	return res

}

func main() {
	
}
