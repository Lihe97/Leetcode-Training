package main

import "fmt"

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {

	arr := make([][10]bool,n)
	res := 0
	for i := 0 ; i < len(reservedSeats) ; i++{
		x := reservedSeats[i][0]
		y := reservedSeats[i][1]
		arr[x-1][y-1] = true
	}
	for i := 0 ;i < len(arr) ; i++{
		res += count(&arr[i])
	}
	return res
}
func count(arr *[10]bool)int{
	res := 0
	temp1 := arr[1] !=true && arr[2]!=true && arr[3]!=true && arr[4]!=true
	temp3 := arr[5] !=true && arr[6]!=true && arr[7]!=true && arr[8]!=true
	temp2 := arr[3] !=true && arr[4]!=true && arr[5]!=true && arr[6]!=true
	if temp1 && temp3 {
		res += 2
	}
	if temp2 && (!temp1&&!temp3){
		res ++
	}
	if temp1 && !temp3{
		res ++
	}
	if !temp1 && temp3{
		res ++
	}

	return res

}
func main() {
	a :=[][]int{{4,3},{1,4},{4,6},{1,7}}
	fmt.Println(maxNumberOfFamilies(4,a))
	
}
