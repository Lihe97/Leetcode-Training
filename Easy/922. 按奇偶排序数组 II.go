package main

import "fmt"

func sortArrayByParityII(A []int) []int {

	i ,j := 0,1
	//if len(A) == 2{
	//	if A[0] % 2 == 1{
	//		A[0],A[1] = A[1],A[0]
	//		return A
	//	}else{
	//		return A
	//	}
	//}

	for (j != len(A)-1 || i != len(A)-2) ||( A[i] % 2 ==1 || A[j] % 2 ==0){
		if A[j] % 2 == 0 || A[i] % 2 == 1{
			A[i],A[j] = A[j],A[i]
		}
		if A[j] % 2 == 1 && j!= len(A)-1{
			j += 2
		}
		if A[i] % 2 == 0 && i!= len(A)-2{
			i += 2
		}

	}

	return A

}

func main() {

	a :=[]int{2,3,5,2}
	fmt.Println(sortArrayByParityII(a))
	
}
