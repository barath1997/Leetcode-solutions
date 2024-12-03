/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 2 pass solution : TC O(2n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    // finding the length O(n)
    length := 0
    curr := head
    for curr != nil {
        curr = curr.Next
        length++
    }

    if length == 1 {
        return nil
    }
    
    // TC : O(n) removing the node
    if length == n {
        return head.Next
    }
    point := length - n
    curr = head
    prev := head
    for point != 0 {
       prev = curr
       curr = curr.Next
       point--
    }

    next := curr.Next
    
    prev.Next = next
    return head
    
}