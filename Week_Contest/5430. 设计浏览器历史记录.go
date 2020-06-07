package main

import "fmt"

type BrowserHistory struct {

	cur int
	temp []string

}


func Constructor(homepage string) BrowserHistory {

	temp := []string{}
	temp = append(temp,homepage)
	return BrowserHistory{
		cur:  0,
		temp: temp,
	}
}
func (this *BrowserHistory) Visit(url string)  {

	if this.cur < len(this.temp){
		this.temp = this.temp[0:this.cur+1]
		this.temp = append(this.temp,url)
	}else{
		this.temp = append(this.temp,url)
	}
	this.cur ++

}

func (this *BrowserHistory) Back(steps int) string {

	if this.cur + 1 <= steps{
		this.cur = 0
		return this.temp[this.cur]

	} else{
		this.cur = this.cur  - steps
		return this.temp[this.cur]
	}
}


func (this *BrowserHistory) Forward(steps int) string {
	if len(this.temp) - 1 - this.cur < steps{
		this.cur = len(this.temp)-1
		return this.temp[len(this.temp)-1]
	} else{
		this.cur += steps
		return this.temp[this.cur]
	}
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
func main() {
	obj := Constructor("a")
	obj.Visit("b")
	obj.Visit("c")
	obj.Visit("d")

	fmt.Println(obj.Back(1))
	fmt.Println(obj.Back(1))
	fmt.Println(obj.Forward(1))
	obj.Visit("e")
	fmt.Println(obj.Forward(2))
	fmt.Println(obj.Back(2))
	fmt.Println(obj.Back(7))




}
