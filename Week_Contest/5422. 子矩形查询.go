package main

import "fmt"

type SubrectangleQueries struct {

	rec [][]int
	rows int
	cols int
}


func Constructor(rectangle [][]int) SubrectangleQueries {

	return SubrectangleQueries{
		rec:  rectangle,
		rows: len(rectangle),
		cols: len(rectangle[0]),
	}
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {

	for i := row1 ; i <= row2 ; i ++{
		for j := col1 ; j <= col2 ; j++{
			this.rec[i][j] = newValue
		}
	}

}


func (this *SubrectangleQueries) GetValue(row int, col int) int {

	return this.rec[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */

func main() {

	rec := [][]int{{1,2,1},{4,3,4}}
	obj := Constructor(rec)
	fmt.Println(obj.rec)
	
}
