# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        def findKthSmallest(root: Optional[TreeNode], k: int) -> tuple[int, int]:
            if root is None:
                return (0,k)
            res, tmp = findKthSmallest(root.left, k)
            if tmp==0:
                return res, tmp
            k = tmp
            if k - 1 == 0:
                return root.val, 0

            return findKthSmallest(root.right, k-1)
        return findKthSmallest(root, k)[0]


             