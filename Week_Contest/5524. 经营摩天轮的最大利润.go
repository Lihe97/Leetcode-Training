package main
func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	res := 0
	max := -1
	ans := -1

	cur := 0

	sum := 0



	for i := 0 ; i < len(customers) ||cur > 0 ; i ++{
		if i < len(customers){
			cur += customers[i]
		}

		if cur >= 4{
			sum += 4
			res = sum * boardingCost - (i+1)*runningCost
			cur -= 4
		}else{
			sum += cur
			res = sum * boardingCost - (i+1)*runningCost
			cur = 0
		}
		if res > max{
			max = res
			ans = i + 1
		}
		//fmt.Println(cur,sum,max)
	}
	return ans

}
func main() {
	
}
