package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {

	res := 0

	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] >= boxTypes[j][1]
	})
	for i := 0 ; i < len(boxTypes) && truckSize > 0  ; i ++{
		if truckSize >= boxTypes[i][0]{
			res += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		}else{

			res += truckSize * boxTypes[i][1]
			truckSize = 0
		}
		//fmt.Println(truckSize)
	}
	return res
}

func main() {
	
}
