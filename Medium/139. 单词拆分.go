package main
func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool,len(s)+1)
	dp[0] = true

	mp := map[string]bool{}
	for _,x := range wordDict{
		mp[x] = true
	}

	for i := 1; i <= len(s) ; i ++{
		for j := 0 ; j < i ; j ++{
			if dp[j] && mp[s[j:i]]{
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	
}
