package main
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	prev := &ListNode{
//		Val:  0,
//		Next: l1,
//	}
//
//	var sum int
//
//	for l1.Next != nil && l2.Next != nil{
//		sum += l1.Val + l2.Val
//		l1.Val = sum % 10
//		sum = sum / 10
//		l1 = l1.Next
//		l2 = l2.Next
//	}
//	sum += l1.Val + l2.Val
//	l1.Val = sum % 10
//	sum = sum / 10
//	if l1.Next == nil && l2.Next == nil && sum == 0{
//		return prev.Next
//	}
//	if l1.Next == nil && l2.Next == nil && sum > 0{
//		node  := &ListNode{
//			Val:  sum,
//			Next: nil,
//		}
//		l1.Next = node
//	}else if l1.Next == nil{
//		l2 = l2.Next
//		l1.Next = l2
//		sum += l2.Val
//		l2.Val = sum % 10
//		sum = sum / 10
//		for sum  > 0 && l2.Next != nil{
//			l2 = l2.Next
//			sum += l2.Val
//			l2.Val = sum % 10
//			sum = sum/10
//		}
//		if sum > 0 {
//			node  := &ListNode{
//				Val:  sum,
//				Next: nil,
//			}
//			l2.Next = node
//		}
//	}else if l2.Next == nil{
//		l1 = l1.Next
//		l2.Next = l1
//		sum += l1.Val
//		l1.Val = sum % 10
//		sum = sum / 10
//		for sum  > 0 && l1.Next != nil{
//			l1 = l1.Next
//			sum += l1.Val
//			l1.Val = sum % 10
//			sum = sum/10
//		}
//		if sum > 0 {
//			node  := &ListNode{
//				Val:  sum,
//				Next: nil,
//			}
//			l1.Next = node
//		}
//	}
//	return prev.Next
//}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	prev := &ListNode{}
	p := prev
	carry := 0
	for l1 != nil || l2 != nil || carry > 0{
		sum := carry

		if l1 != nil{
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			sum += l2.Val
			l2 = l2.Next
		}
		node := &ListNode{
			Val:  sum%10,
			Next: nil,
		}
		p.Next = node
		p = p.Next
		carry = sum/10
	}
	return prev.Next
}
func main() {
	
}
