/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    
    l := 0
    mid := 0
    temp := head
    for temp != nil {
        temp = temp.Next
        l++
    }

       mid = (l/2) + 1
       
       temp2 := head
       count := 0
       for count != mid-1 {
          temp2 = temp2.Next
          count++
       }

       return temp2

}