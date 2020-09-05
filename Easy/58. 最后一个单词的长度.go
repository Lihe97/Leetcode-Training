package main

func lengthOfLastWord(s string) int {

	l := 0

	for i := len(s) - 1 ; i >= 0 ; i--{
		if s[i] != ' '{
			l ++
			continue
		}else if s[i] == ' ' && l == 0{
			continue
		}
		break
	}
	return l

}
func main() {
	
}
