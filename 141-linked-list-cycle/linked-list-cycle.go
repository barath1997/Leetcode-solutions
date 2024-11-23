/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // TC : O(N) , SC : O(1)
/*func hasCycle(head *ListNode) bool {
    
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
}*/

// TC : O(N) , SC : O(N) using map
func hasCycle(head *ListNode) bool {

    m := make(map[*ListNode]bool)

    temp := head

    for temp != nil {
        if val := m[temp];val {
            return true
        } else {
           m[temp] = true
        }
       
       temp = temp.Next
    }
    return false
}
