package main

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	isLessThanZero := false
	i := 0
	for str[i] == ' ' {
		i ++
		if i >= len(str) {
			return 0
		}
	}
	str = str[i:]
	if str[0] == '+' {
		str = str[1:]
	} else if str[0] == '-' {
		str = str[1:]
		isLessThanZero = true
	}
	var res uint64 // 便于检查溢出
	i = 0
	for {
		if res >= 2147483647 && !isLessThanZero{
			return 2147483647
		} else if res >= 2147483648 && isLessThanZero {
			return -2147483648
		}
		if i >= len(str) {
			break
		}
		if str[i] - '0' < 0 ||
			str[i] - '0' > 9 {
			if isLessThanZero {
				return -1*int(res)
			} else {return int(res)}
		}
		res = res * 10 + uint64(str[i] - '0')
		i ++
	}
	if isLessThanZero {
		return -1*int(res)
	}
	return int(res)
}

func main() {
	
}
