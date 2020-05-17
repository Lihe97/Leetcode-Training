package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {


	res := 0

	for i := 0 ; i < len(startTime) ; i++{
		if queryTime >=startTime[i] && queryTime <= endTime[i]{
			res ++
		}
	}
	return res

}

func main() {
	
}
