package main
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	cura := headA
	curb := headB
	for cura != curb{
		if cura == nil{
			cura = headB
		}else{
			cura = cura.Next
		}

		if curb == nil{
			curb = headA
		}else{
			curb = curb.Next
		}
	}
	return cura
}
func main() {
	
}
