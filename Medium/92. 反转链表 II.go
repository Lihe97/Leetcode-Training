package main

import (
	"fmt"
)

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	temp := head
	start := head
	end := head
	for i := 0 ; i< n;i++{
		end = end.Next
	}
	for i := 1 ; i< m-1;i++{
		start = start.Next
	}
	for  i := 1 ; i < m ; i++{
		temp = temp.Next
	}
	//fmt.Println(end)
	//fmt.Println("end",end)
	//fmt.Println("temp",temp)
	for i := 0 ; i <= n - m; i++{
		next := temp.Next
		temp.Next = end
		end = temp
		temp = next
	}
	//fmt.Println()
	//fmt.Println("end",end)
	//fmt.Println("temp",temp)
	//fmt.Println("head",head)


	if m == 1{
		head = end
		return head
	}
	start.Next = end

	//for head!=nil{
	//	fmt.Println(head)
	//	head = head.Next
	//	time.Sleep(2*time.Second)
	//}
	return head


}
func main() {
	//f := ListNode{
	//	Val:  6,
	//	Next: nil,
	//}
	e := ListNode{
		Val:  5,
		Next: nil,
	}
	d:= ListNode{
		Val:  4,
		Next: &e,
	}
	c:= ListNode{
		Val:  3,
		Next: &d,
	}
	b := ListNode{
		Val:  2,
		Next: &c,
	}
	a:= ListNode{
		Val:  1,
		Next: &b,
	}
	fmt.Println(reverseBetween(&a,2,4))
	
}
