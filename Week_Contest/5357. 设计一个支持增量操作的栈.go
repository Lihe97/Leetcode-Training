package main

import "fmt"

type CustomStack struct {
	values    []int
	top int

}
func Constructor(maxSize int) CustomStack {

	return CustomStack{
		values: make([]int,maxSize),
		top:    -1,
	}
}
func (this *CustomStack) Push(x int)  {
	if this.top == len(this.values)-1{
		return
	}
	this.top ++
	this.values[this.top] = x
}


func (this *CustomStack) Pop() int {
	if this.top == -1 {
		return -1
	}
	e := this.values[this.top]
	this.top --
	return e

}


func (this *CustomStack) Increment(k int, val int)  {
	length := 0
	for i:= 0 ;i <= this.top;i++{
		length ++
	}
	temp := 0
	if k >= length{
		temp = length
	}else{
		temp = k
	}
	for j:= 0 ; j <temp;j++{
		this.values[j] += val
	}

}
func main() {
	 obj := Constructor(3)
	 obj.Push(1)
	 obj.Push(2)
	 obj.Push(3)

	 obj.Increment(3,10)
	param_2 := obj.Pop()
	 fmt.Println(param_2)
}
