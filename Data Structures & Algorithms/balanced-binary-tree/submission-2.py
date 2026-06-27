# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:

        def isBalanced(node: Optional[TreeNode], height: int) -> (int, bool):
            if not node:
                return (height, True)
            
            leftHeight, leftBalanced = isBalanced(node.left, height)
            rightHeight, rightBalanced = isBalanced(node.right, height)

            balanced = leftBalanced and rightBalanced and abs(leftHeight - rightHeight)<=1
            return (max(leftHeight, rightHeight)+1, balanced)

        return isBalanced(root, 0)[1]
        