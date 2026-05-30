/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list2 == nil{
        return list1
    }
    if list1 == nil {
        return list2
    }

    var c *ListNode
    if list1.Val < list2.Val {
        c = list1
        list1 = list1.Next
    }else{
        c = list2
        list2 = list2.Next
    }
        h := c
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            c.Next = list1
            list1 = list1.Next
        }else{
            c.Next = list2
            list2 = list2.Next
        }
        c = c.Next
    }
    for list1 != nil {
        c.Next = list1
        c = c.Next
        list1 = list1.Next
    }
    for list2 != nil {
        c.Next = list2
        c = c.Next
        list2 = list2.Next
    }
    return h
}
