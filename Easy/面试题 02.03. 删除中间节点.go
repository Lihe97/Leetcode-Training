package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//sb题
func ddeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next

}
func main() {
	
}
