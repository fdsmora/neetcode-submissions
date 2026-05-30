/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    n := getLength(head)

    half := n/2

    c := head

    for i:=0; i<half; i++ {
        c = c.Next
    }

    tmp := c.Next
    c.Next = nil
    c = tmp

  r := reverseList(c)

    l := head

    dummy := &ListNode{}
    c = dummy

    for l != nil && r != nil {
        c.Next = l
        l = l.Next
        c = c.Next
        c.Next = r
        r = r.Next
        c = c.Next
    }

    for l != nil {
        c.Next = l
        c = c.Next
        l = l.Next
    }
    for r != nil {
        c.Next = r
        c = c.Next
        r = r.Next
    }
    head = dummy.Next
}

func getLength(head *ListNode) int {
    c := head
    count := 0
    for c != nil {
        count++
        c = c.Next
    }
    return count
}

func reverseList(c *ListNode) *ListNode{
    var prev *ListNode 
    for c != nil {
        tmp := c.Next
        c.Next = prev
        prev = c 
        c = tmp
    }
    return prev
}
