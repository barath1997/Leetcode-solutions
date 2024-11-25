/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // Optimised TC : O(N/2) , SC : O(1)
 /*func reverseList(head *ListNode) *ListNode {

    if head == nil {return head}

    var prev *ListNode = nil
    var next *ListNode = nil
    var curr *ListNode = head

    for curr != nil {
        next = curr.Next // Store the Next Node before Changing curr 
        curr.Next = prev // Actual Reversal
        
        // Move prev and next by 1 step
        prev = curr
        curr = next
    }

    head = prev
    return head
 }*/

 // TC : O(2N) , SC : O(N)  stack based solution
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {return head}
 
	curr := head
	
	var st Stack
 
	// inserting nodes into stack : O(n)
    for curr != nil {
        st.Push(curr)
        curr = curr.Next
    }

    // creating a new linkedList
    res := &ListNode{Val:-1,Next:nil}
    curr = res
    
    // till stack is not empty, pop elements and put inside a linkedlist
    for !st.IsEmpty() {
        curr.Next = st.Pop()
        curr = curr.Next
    }
    
    // curr is now at last node, now last node should be directing to nil
    curr.Next = nil
    
    // since res has first element as -1 so we have to return res.Next, which has all the elements in reverse.
    return res.Next

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
 
 }