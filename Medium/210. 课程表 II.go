package main

func findOrder(numCourses int, prerequisites [][]int) []int {

	res := []int{}
	queue := []int{}
	table := make([]int,numCourses)

	for i := 0 ; i < len(prerequisites); i++{
		table[prerequisites[i][0]] ++
	}
	for i := 0 ; i < len(table) ; i++{
		if table[i] == 0{
			queue = append(queue,i)
		}
	}
	for len(queue) != 0{
		temp := queue[0]
		queue = queue[1:]
		res = append(res,temp)

		for i := 0 ; i < len(prerequisites) ; i++{
			if prerequisites[i][1] == temp{
				table[prerequisites[i][0]] --
				if table[prerequisites[i][0]] == 0 {
					queue = append(queue,prerequisites[i][0])
				}
			}
		}
	}
	if len(res) == numCourses{
		return res
	}else{
		return nil
	}
}

func main() {
	
}
