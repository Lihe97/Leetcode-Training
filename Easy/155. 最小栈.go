package main
type MinStack struct {
	values []int
	min []int
}


/** initialize your data structure here. */
func CConstructor() MinStack {
	return MinStack{
		values: make([]int,0),
		min:    make([]int,0),
	}
}


func (this *MinStack) Push(x int)  {

	this.values = append(this.values, x)
	if len(this.min) == 0{
		this.min = append(this.min, x)
	}else{
		if this.min[len(this.min) - 1] > x{
			this.min = append(this.min, x)
		}else{
			this.min = append(this.min, this.min[len(this.min)-1])
		}
	}
}


func (this *MinStack) Pop()  {

	this.values = this.values[0:len(this.values)-1]
	this.min = this.min[0:len(this.min)-1]

}


func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]

}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]

}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := CConstructor()
	obj.Push(2)
	obj.GetMin()
	
}
