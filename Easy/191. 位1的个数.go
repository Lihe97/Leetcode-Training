package main
//func hammingWeight(num uint32) int {
//	var mask uint32
//	mask = 1
//	count := 0
//	for i := 0 ; i < 32 ; i ++{
//
//		if (num & mask) != 0{
//			count ++
//		}
//		mask = mask << 1
//	}
//	return count
//}
func hammingWeight(num uint32) int {

	count := 0
	for num != 0 {
		num &= num-1
		count ++
	}

	return count
}

func main() {
	
}
