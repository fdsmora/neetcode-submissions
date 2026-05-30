/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    Q := []*TreeNode{root}

    for len(Q)>0 {
        c := Q[0]
        Q = Q[1:]
        if c!=nil {
            Q = append(Q, c.Left)
            Q = append(Q, c.Right)
        }
        if checkEqual(c, subRoot) {
            return true
        }
    }
    return false
}

func checkEqual(root, subRoot *TreeNode) bool {
    if (root == nil && subRoot == nil) {
        return true
    }

    return root != nil && subRoot != nil && 
        root.Val == subRoot.Val &&
        checkEqual(root.Left, subRoot.Left) && 
        checkEqual(root.Right, subRoot.Right)
}

/*
Q = {1,2,3}
c = 1
sr = 2

*/
