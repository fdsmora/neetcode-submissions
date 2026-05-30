/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for slow.Next != nil {
        slow = slow.Next
        if fast.Next != nil {
            fast = fast.Next
            if fast.Next != nil {
                fast = fast.Next
                if fast == slow {
                    return true
                }
            }else{
                return false
            }
        }else {
            return false
        }
    }
    return false
}
