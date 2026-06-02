# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        res = 0
        count = k

        def dfs(root: Optional[TreeNode]):
            nonlocal res, count
            if root is None:
                return
            dfs(root.left)
            count-=1
            if count == 0:
                res = root.val
                return
            return dfs(root.right)

        dfs(root)
        return res

        #return kthSmallest(self,root, k)
    # def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
    #     return self.findKthSmallest(root, k)[0]
    # def findKthSmallest(self, root: Optional[TreeNode], k: int) -> tuple[int, int]:
    #     if root is None:
    #         return (0,k)
    #     res, tmp = self.findKthSmallest(root.left, k)
    #     if tmp==0:
    #         return res, tmp
    #     k = tmp
    #     if k - 1 == 0:
    #         return root.val, 0

    #     return self.findKthSmallest(root.right, k-1)

             