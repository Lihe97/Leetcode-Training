package main
func rotateString(A string, B string) bool {
	if A==B {
		return true
	}

	for i := 0; i < len(A); i++ {
		if A[i:] + A[:i] == B {
			return true
		}
	}
	return false
}
func main() {
	
}
