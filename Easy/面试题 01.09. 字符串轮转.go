package main
func isFlipedString(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	mp := map[byte]int{}
	for i := 0 ; i < len(s1) ; i ++{
		mp[s1[i]] ++
		mp[s2[i]] --
	}
	for _, a := range mp{
		if a != 0{
			return false
		}
	}
	return true
}

func main() {
	
}
