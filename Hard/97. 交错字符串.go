package main
func isInterleave(s1 string, s2 string, s3 string) bool {

	if len(s3) != len(s1) + len(s2) {
		return false
	}

	f := make([][]bool,len(s1) + 1)

	for i := 0 ; i <= len(s1) ; i ++{
		f[i] = make([]bool,len(s2)+1)
	}

	f[0][0] = true

	for i := 0 ; i <= len(s1) ; i ++{
		for j := 0 ; j <= len(s2) ; j++{
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}

	return f[len(s1)][len(s2)]

}

func main() {
	
}
