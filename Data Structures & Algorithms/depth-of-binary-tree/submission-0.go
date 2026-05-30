/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxLeft := calcMaxDepth(root.Left, 1)
    maxRight := calcMaxDepth(root.Right, 1)
    return max(maxLeft, maxRight)+1
}

func calcMaxDepth(root *TreeNode, depth int) int {
    if root == nil {
        return depth -1
    }
    maxLeft := calcMaxDepth(root.Left, depth + 1)
    maxRight := calcMaxDepth(root.Right, depth + 1)
    return max(maxLeft, maxRight)
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
