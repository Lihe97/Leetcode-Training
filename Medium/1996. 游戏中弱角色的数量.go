package main

import "sort"

func numberOfWeakCharacters(p [][]int) int {

	sort.Slice(p, func(i, j int) bool {
		a, b := p[i], p[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	max := -1
	cnt := 0
	for _, v := range p {
		if v[1] < max {
			cnt++
		} else {
			max = v[1]
		}
	}
	return cnt
}
