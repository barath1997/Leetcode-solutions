/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // Optimised TC : O(N/2) , SC : O(1)
 func reverseList(head *ListNode) *ListNode {

    if head == nil {return head}

    var prev *ListNode = nil
    var next *ListNode = nil
    var curr *ListNode = head

    for curr != nil {
        next = curr.Next
        curr.Next = prev
        
        prev = curr
        curr = next
    }

    head = prev
    return head
 }

 // TC : O(2N) , SC : O(N)  stack based solution
/*func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {return head}
 
	temp := head
	
	var st Stack
 
	// inserting nodes into stack : O(n)
	for temp != nil {
	  st.Push(temp)
	  temp = temp.Next
	} 
 
	ll := &ListNode{Val: -1,Next:nil}
    temp = ll
	for !st.IsEmpty() {
       
       temp.Next = st.Pop()
       temp = temp.Next
	}
	
	temp.Next = nil
	return ll.Next
 }
 
 type Stack []*ListNode
 
 func (s *Stack) Push(n *ListNode) {
	 *s = append(*s,n)
 }
 
 func (s *Stack) Pop() *ListNode{
	 temp := *s
	 val := temp[len(*s)-1]
	 *s = temp[:len(*s)-1]
	 return val
 }
 
 func (s *Stack) IsEmpty() bool{
	 return len(*s) == 0
 }
 
 func (s *Stack) Peek() *ListNode{
	 
	 if !s.IsEmpty() {
		 temp := *s
		 return temp[len(*s)-1]
	 }
	 return nil
 
 }*/