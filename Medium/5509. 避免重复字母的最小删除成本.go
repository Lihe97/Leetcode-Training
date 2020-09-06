package main

import "sort"

func minCost(s string, cost []int) int {

	res := 0

	temp := []int{}
	for i := 0 ; i < len(s) -1 ; {
		j := i + 1
		for j < len(s) && s[j] == s[i]  {
			temp = append(temp,cost[j])
			//fmt.Println(temp,i,j)

			j ++
		}

		if len(temp) != 0{
			temp = append(temp,cost[i])
		}
		i = j
		//fmt.Println(temp)
		sort.Ints(temp)
		for k := 0 ; k < len(temp) - 1; k ++{
			res += temp[k]
		}
		temp = []int{}
	}




	return res

}

func main() {
	
}
