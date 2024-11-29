/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// TC : O(left) + O(right-left)
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	dummy := &ListNode{Val: -1, Next: nil}
	dummy.Next = head
	prev := dummy
	curr := dummy.Next

	i := 1
	for i < left {
		prev = curr
		curr = curr.Next
		i++
	}

	res := prev
	for i <= right {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
		i++
	}

    res.Next.Next = curr

    res.Next = prev

    return dummy.Next

}