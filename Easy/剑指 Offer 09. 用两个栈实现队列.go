package main
type CQueue struct {

	s1 []int
	s2 []int
}
func Constructor() CQueue {

	c1 := []int{}
	c2 := []int{}
	return CQueue{
		s1: c1,
		s2: c2,
	}
}


func (this *CQueue) AppendTail(value int)  {

	this.s1 = append(this.s1,value)

}


func (this *CQueue) DeleteHead() int {

	if len(this.s1) == 0 && len(this.s2) == 0{
		return -1
	}

	if len(this.s2) == 0{
		for len(this.s1) > 0{
			temp := this.s1[len(this.s1)-1]
			this.s2 = append(this.s2,temp)
			this.s1 = this.s1[:len(this.s1)-1]

		}
	}

	temp := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return temp

}
func main() {
	
}
