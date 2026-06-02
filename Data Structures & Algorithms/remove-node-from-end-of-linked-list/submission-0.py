# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        count = 0
        current = head
        pre = ListNode(0,head)
        while current:
            count+=1
            current = current.next
        target = count - n
        current = pre
        while target:
            current = current.next
            target-=1
        if current.next:
            tmp = current.next
            current.next = current.next.next
            tmp.next = None
        else:
            current.next = None
        return pre.next
