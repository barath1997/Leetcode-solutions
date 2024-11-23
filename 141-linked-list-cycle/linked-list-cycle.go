/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    
    const maxLen = 10000

    temp := head
    iterations := 0
    for temp != nil {
        iterations++
        if iterations > maxLen {
          return true
        }
        temp = temp.Next
    }

    return false
}