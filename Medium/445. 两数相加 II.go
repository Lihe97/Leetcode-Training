package main

type ListNode struct {
	Val int
	Next *ListNode
}
func aaddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	stack1 := []int{}
	stack2 := []int{}
	for l1 != nil{
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil{
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	jw := 0
	prev := &ListNode{
		Val:  0,
		Next: nil,
	}
	for len(stack1) != 0 || len(stack2) != 0 || jw != 0{
		sum := jw
		if len(stack1) != 0{
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0{
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		node := &ListNode{
			Val:  sum%10,
			Next: prev.Next,
		}
		prev.Next = node
		jw = sum / 10
	}
	return prev.Next
}
func main() {
	
}
