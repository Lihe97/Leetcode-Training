package main

import "fmt"

func isPalindrome(head *ListNode) bool {

	res := []int{}
	for head != nil{
		res = append(res, head.Val)
		head = head.Next
	}

	for i:= 0 ; i < len(res)/2 ; i ++{
		if res[i] != res[len(res)-i-1]{
			return false

		}

	}

	return true
}

func main() {
	e := ListNode{
		Val:  1,
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
	fmt.Println(isPalindrome(&a))
	
}
