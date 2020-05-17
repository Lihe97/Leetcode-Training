package main

func canFinish(numCourses int, prerequisites [][]int) bool {


	queue := []int{}
	table := make([]int,numCourses)

	for i := 0 ; i < len(prerequisites) ; i ++{
		table[prerequisites[i][0]] ++
	}
	res := []int{}
	for a,b := range table{
		if b == 0{
			queue = append(queue,a)
		}
	}

	for len(queue) != 0{
		temp := queue[0]
		queue = queue[1:]
		res  = append(res,temp)
		for i := 0 ; i < len(prerequisites) ; i ++{
			if prerequisites[i][1] == temp{
				table[prerequisites[i][0]] --
				if table[prerequisites[i][0]] == 0{
					queue = append(queue,prerequisites[i][0])
				}

			}
		}
	}
	return len(res) == numCourses
}

func main() {
	
}
