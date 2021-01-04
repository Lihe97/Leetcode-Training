package main

import "sort"

func lastStoneWeight(stones []int) int {



	sort.Ints(stones)


	for len(stones) != 1{
		last := stones[len(stones) - 1]
		llast := stones[len(stones) - 2]
		if last == llast{
			if len(stones) == 2{
				return 0
			}
			stones = stones[:len(stones) - 2]
		}else{
			stones = stones[:len(stones)-1]
			stones[len(stones) - 1] = last - llast
		}
		sort.Ints(stones)
	}
	return stones[0]

}

func main() {
	
}
