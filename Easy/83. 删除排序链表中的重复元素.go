package main


func deleteDuplicates(head *ListNode) *ListNode {

	prev := &ListNode{0, head}
	cur := prev
	count := 0
	for(head != nil && head.Next != nil) {
		for (head.Next != nil && head.Val == head.Next.Val ){
			count++
			head=head.Next
		}
		if count > 0 {
			cur.Next = head.Next
			head=head.Next
			count = 0
		}else{
			cur.Next = head
			cur = cur.Next
			head = head.Next
		}
	}
	return prev.Next
}
func main() {


}
