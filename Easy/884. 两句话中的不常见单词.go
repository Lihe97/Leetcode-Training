package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {

	mp1 := map[string]int{}
	mp2 := map[string]int{}

	str1 := strings.Fields(s1)
	str2 := strings.Fields(s2)

	for _, a := range str1 {
		mp1[a]++
	}
	for _, a := range str2 {
		mp2[a]++
	}
	res := []string{}
	for k, v := range mp1 {
		if mp2[k] == 0 && v == 1 {
			res = append(res, k)
		}
	}
	for k, v := range mp2 {
		if mp1[k] == 0 && v == 1 {
			res = append(res, k)
		}
	}
	return res
}
