package main

import "sort"

func maxEvents(events [][]int) int {

	eLen := len(events)
	if eLen <= 1 {
		return eLen
	}

	sort.Slice(events, func(i int, j int) bool {
		if events[i][1] != events[j][1]{
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	// fmt.Println(minDay)

	days := make([]bool, 100005)
	maxEvent := 0
	for i := 0; i < eLen; i++ {
		for j := events[i][0]; j <= events[i][1]; j++ {
			if !days[j] {
				days[j] = true
				maxEvent++
				break
			}
		}
	}

	return maxEvent

}
func main() {
	
}
