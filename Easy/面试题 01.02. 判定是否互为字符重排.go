package main
func CheckPermutation(s1 string, s2 string) bool {
	hh := make(map[byte]int)
	if len(s1) != len(s2){
		return false
	}
	for i := 0;  i < len(s1); i++{
		hh[s1[i]] ++
		hh[s2[i]] --

	}
	for _,v  := range hh{
		if v!= 0{
			return false
		}

	}
	return true


}
func main() {
	
}
