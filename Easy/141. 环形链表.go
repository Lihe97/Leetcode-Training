package main

//func hasCycle(head *ListNode) bool {
//
//	mp := map[ListNode]int{}
//	for head != nil{
//		mp[*head] ++
//		if mp[*head] > 1{
//			return false
//		}
//		head = head.Next
//	}
//	return true
//}
func hasCycle(head *ListNode) bool {

	if head == nil{
		return false
	}
	first := head.Next
	second := head
	for first !=nil && second != nil && first.Next != nil{
		if first == second{
			return true
		}
		first = first.Next.Next
		second = second.Next
	}
	return false
}



func main() {
	
}
