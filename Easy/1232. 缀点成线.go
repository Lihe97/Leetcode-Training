package main
func checkStraightLine(coordinates [][]int) bool {

	y := coordinates[1][1] - coordinates[0][1]
	x := coordinates[1][0] - coordinates[0][0]
	for i:=2 ; i < len(coordinates); i++{
		y1 := coordinates[i][1] - coordinates[i-1][1]
		x1 := coordinates[i][0] - coordinates[i-1][0]
		if y*x1 != x*y1{
			return false
		}
	}
	return true


}
func main() {
	
}
