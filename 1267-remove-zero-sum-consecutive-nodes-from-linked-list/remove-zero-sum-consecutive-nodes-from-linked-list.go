/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// TC : O(N) , SC : O(N) refer "https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/solutions/366350/c-o-n-explained-with-pictures/"

func removeZeroSumSublists(head *ListNode) *ListNode {

	prefixLinkedListMap := make(map[int]*ListNode)

	root := &ListNode{Val: 0, Next: head}
    
    // set the 0th accumulatedSum to root Node, so that when we see a 0 repeated this will helpful
    prefixLinkedListMap[0] = root

	accumulatedSum := 0

	for head != nil {
        
        // adding head.Val to "accumulatedSum" to get prefixSum
		accumulatedSum += head.Val

		// check if we have already seen the sum
		if prev, exists := prefixLinkedListMap[accumulatedSum]; exists {
            
            // start is set so that we can skip the useless nodes and set start.Next = head.Next
			start := prev.Next
            
            // sum is used to get accumulated sum from prev to curr head (where we find the repeated value)
			sum := accumulatedSum

            // iterate over start pointer and delete the useless key,value pairs in the map
			for start != head {
				sum += start.Val
				delete(prefixLinkedListMap, sum)
                start = start.Next
				
			}

			// remove all the nodes that are not contrubuting to the sum as previous and curr values are same so
			// its a zero sum game , so remove them
			prev.Next = head.Next

		} else {
			// if not exists then add this value to map
			prefixLinkedListMap[accumulatedSum] = head
		}

		head = head.Next

	}

    return root.Next

}