package main
type point struct {
	x int
	y int
}
func isPathCrossing(path string) bool {

	mp := map[point]bool{}

	curx ,cury := 0,0
	mp[point{curx,cury}] = true
	for i := 0 ; i < len(path) ; i++{
		if path[i] == 'N'{
			cury ++
		}
		if path[i] == 'S'{
			cury --
		}
		if path[i] == 'E'{
			curx ++
		}
		if path[i] == 'W'{
			curx --
		}
		if mp[point{curx,cury}]{
			return true
		}else{
			mp[point{curx,cury}] = true
		}
	}
	return false

}

func main() {
	
}
