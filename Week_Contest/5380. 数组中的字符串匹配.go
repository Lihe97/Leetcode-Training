package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {

	res := []string{}
	mp := map[string]bool{}
	for _,x := range words{
		mp[x] = false
	}

	for i := 0 ; i < len(words) ; i ++{
		if mp[words[i]] == true{
			continue
		}else{
			for j := i+1; j < len(words) && mp[words[i]] == false; j++ {

				if strings.Contains(words[i], words[j]) && mp[words[j]] == false {
					res = append(res, words[j])
					mp[words[j]] = true
				}
				if strings.Contains(words[j], words[i]) {
					res = append(res, words[i])
					mp[words[i]] = true
					break
				}
			}
		}
	}

	return res
}

func main() {

	a := []string{"kohdqtlhtwhxfw","acftf","kohdqtlhtwhxodp","cwpeiowwms","ehvjk","mkdmsmxb","mehvjks","mxdoz","xnacftf","mgksgencwhk","hacftf","jdofko","mcwpeiowwms","x","pommgeuefd"}
	fmt.Println(stringMatching(a))
}
