package main

func plusOne(digits []int) []int {

	flag := 1

	for i := len(digits) - 1; flag != 0 && i >= 0 ; i --{

		digits[i] += flag
		flag = digits[i]/10
		digits[i] = digits[i]%10

	}
	if flag == 1{
		digits = append([]int{1},digits...)
	}
	return digits
}

func main() {
	
}
