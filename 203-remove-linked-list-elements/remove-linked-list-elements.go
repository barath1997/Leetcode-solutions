/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

    if head == nil {return head}
    if  head.Next == nil {
        if head.Val == val { head = head.Next; return head } else {
            return head
        }
    }

    temp := head

    for temp != nil {
        if temp.Val == val {
            head = head.Next
            temp = head
            continue
        } else if temp.Next != nil && temp.Next.Val == val {
            temp.Next = temp.Next.Next
            continue
        }
        temp = temp.Next
    }

    return head
    
}