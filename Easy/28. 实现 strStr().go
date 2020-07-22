package main

func strStr(haystack string, needle string) int {

	res := -1
	if len(needle) == 0{
		return 0
	}

	n := 0

	for i := 0 ; i < len(haystack) ; i ++{
		m := i
		for m < len(haystack) && n < len(needle) && haystack[m] == needle[n]  {
			m ++
			n ++
		}
		if n == len(needle){
			return i
		}else{
			n = 0
		}

	}


	return res

}
func main() {
	
}
