/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	// finding the middle node

	if head.Next == nil {
		return
	}

	curr := head
	len := 0
	var fast, slow *ListNode
	fast = curr
	slow = curr
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		println("slow : ", slow.Val)
	}

	l := head

	for l != nil {
		l = l.Next
		len++
	}

	if len%2 != 0 {
		len++
	}

	slow = slow.Next

	// reversing the right half of linkedlist
	var prev, next *ListNode = nil, nil

	for slow != nil {
		next = slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	head1, head2 := head, prev
	// intertwine the linkedlist now
	for head2 != nil {

		next := head1.Next
		head1.Next = head2
		head1 = head2
		head2 = next

	}

	temp := head
	for len > 0 {
		temp = temp.Next
		len--
	}

	temp.Next = nil
}