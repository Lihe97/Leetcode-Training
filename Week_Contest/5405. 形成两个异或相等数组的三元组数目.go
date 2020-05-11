package main
func countTriplets(arr []int) int {

	res := 0

	for i := 0 ; i < len(arr)-1 ; i ++{

		a := arr[i]

		for j := i + 1; j < len(arr) ; j ++{

			b := arr[j]

			for k := j ; k  < len(arr) ; k++{

				if k > j{
					b ^= arr[k]
				}

				if a == b{
					res ++
				}

			}
			a ^= arr[j]


		}
	}





	return res
}

func main() {
	
}
