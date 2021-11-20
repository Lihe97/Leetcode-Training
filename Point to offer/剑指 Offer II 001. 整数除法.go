package Point_to_offer


func divide(a int, b int) int {

	if a == 0{
		return 0
	}
	if a == b{
		return 1
	}
	if a < b{

	}
	res := 0
	flag := 1
	if a * b < 0{
		flag = -1
	}


	if a < 0{
		a = -a
	}
	if b < 0{
		b = -b
	}
	for a > b{
		a -= b
		res ++
	}

	return res * flag
}