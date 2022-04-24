package main

import "strings"

func toGoatLatin(sentence string) string {

	mp := map[byte]bool{}
	mp['a'] = true
	mp['e'] = true
	mp['i'] = true
	mp['o'] = true
	mp['u'] = true
	mp['A'] = true
	mp['E'] = true
	mp['I'] = true
	mp['O'] = true
	mp['U'] = true
	res := ""
	arr := strings.Split(sentence, " ")

	for i := 0; i < len(arr); i++ {
		var temp string
		if _, ok := mp[arr[i][0]]; ok {
			temp = arr[i] + "ma"
		} else {
			temp = arr[i][1:] + string(arr[i][0])
			temp += "ma"
		}
		for j := 0; j < i+1; j++ {
			temp += "a"
		}
		res += temp
		if i != len(arr)-1 {
			res += " "
		}
	}
	return res
}
