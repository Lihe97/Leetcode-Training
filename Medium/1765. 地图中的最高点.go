package main

type pair struct{ x, y int }

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	q := []pair{}
	for i, row := range isWater {
		for j, water := range row {
			if water == 1 {
				ans[i][j] = 0
				q = append(q, pair{i, j})
			}
		}
	}
	var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && ans[x][y] == -1 {
				ans[x][y] = ans[p.x][p.y] + 1
				q = append(q, pair{x, y})
			}
		}
	}
	return ans

}
