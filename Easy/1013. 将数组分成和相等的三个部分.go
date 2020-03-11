package main

import "fmt"


func canThreePartsEqualSum(A []int) bool {
	sum := 0

	for i:= 0 ; i < len(A) ; i++{
		sum += A[i]
	}
	if sum % 3 != 0{
		return false
	}
	i,j := 0 ,len(A) - 1
	num := sum/3
	sum1,sum2 := 0,0
	for  ; i < len(A) ; i++{
		sum1 += A[i]
		//fmt.Println(sum1, i )
		if sum1 == num{
			break
		}
	}
	//fmt.Println(i)
	//fmt.Println("=--------------=")
	for  ; j >0 ; j--{
		sum2 += A[j]
		//fmt.Println(sum2, j )
		if sum2 == num{
			break
		}
	}
	fmt.Println(i,j)
	if sum1 == sum2 && sum1 == 0 && num == 0 && i +1 < j{
		return true
	}
	if sum1 == sum || sum2 ==sum{
		return false
	}
	//fmt.Println(num)
	fmt.Println(i,j)
	//fmt.Println(sum1,sum2)
	if i< j{
		return true
	}else{
		return false
	}


}
func main() {
	a := []int{1,-1,1,-1}
	fmt.Println(canThreePartsEqualSum(a))
}
