package main

func minOperations(logs []string) int {

	res := 0

	for i := 0 ; i < len(logs) ; i ++{
		if logs[i] == "../"{
			if res == 0{
				continue
			}else{
				res --
			}

		}else if logs[i] == "./"{
			continue
		}else{
			res ++
		}
	}
	return res
}
func main() {
	
}
