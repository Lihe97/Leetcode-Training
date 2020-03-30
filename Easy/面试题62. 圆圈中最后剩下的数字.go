package main

import (
	"fmt"
)

type CycleList struct {
	val int
	next *CycleList
}
func llastRemaining(n int, m int) int {
	if m == 1{
		return n-1
	}
	head := &CycleList{0,nil}
	cur := head
	for i := 1 ; i < n ; i++{
		node := &CycleList{i,nil}
		head.next = node
		head = node
	}
	head.next = cur

	for {

		for i:= 0 ; i < m-2 ; i ++{
			cur = cur.next
		}
		if cur.next.next == cur {
			return cur.val
		}else{
			cur.next = cur.next.next
			cur = cur.next
		}

	}


}
func lastRemaining(n int, m int) int {
	f := 0
	for i := 2; i <= n; i++ {
		f = (m + f) % i
	}
	return f+1
}
func main() {

	fmt.Println(lastRemaining(4,2))
	
}
