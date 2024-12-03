/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

// discussed in class
// TC: O(N)
// SC: O(1) - Extra Space
// 	O(N) - Recursion Stack/ Auxiliary Space 
class Solution {
    public ListNode reverseKGroup(ListNode head, int K) 
    {
        if(head == null)
            return head;
        
        ListNode curr = head;
        int count = 0;

        // Reverse first K Nodes of Linked List
        while (curr!=null && count<K) // O(K)
        {
            curr = curr.next;
            count++;
        }

    /*
        After this while loop, curr.next ---> (K+1)th Node
        2 Cases:
        (1) K = N: Nothing left to reverse
        (2) K!= N: Reverse for Pending Node in LL
    */

        if (count == K)
        {
            // Recur starting from (K+1)th Node
            curr = reverseKGroup(curr, K); // O(N-K)

            // head ---> pointer to direct part
            // curr ---> pointer to revrese part
            // count = k: Reversing K Times

            while (count-- > 0) // count = K, initially
            {
                ListNode temp = head.next;
                head.next = curr;
                curr = head;
                head = temp;
            }

            head = curr; // Reverse
        }

    // No Else Case ----> Directly return head ----> Return as it as when window_Size < k -> No reversals
        return head;
    }
}