# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        
        def DFS(root: Optional[TreeNode], total: int) -> bool:
            if not root:
                return False
            if not root.left and not root.right:
                return root.val+total == targetSum
            return DFS(root.left, total+root.val) or DFS(root.right, total+root.val)

            # if not root:
            #     #print(f"leave, total is {total} and targetSum is {targetSum}")
            #     return total == targetSum
            
            # #print(f"node:{root.val}, total received is {total}")
            # if DFS(root.left, total+root.val):
            #     return True
            
            # return DFS(root.right, total+root.val)

        return DFS(root, 0)