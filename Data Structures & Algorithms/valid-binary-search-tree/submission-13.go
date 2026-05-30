/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
func isValidBST(node *TreeNode) bool {
	return visit(node.Left, node.Val, 0) && visit(node.Right, node.Val, 1)
}

func visit(node *TreeNode, parent int, side int)bool{
	if node == nil {
		return true
	}
	if side == 0 {
		return parent>node.Val && visit(node.Left, node.Val, 0) && visit(node.Right, node.Val, 1)
	}
	return parent<node.Val && visit(node.Left, node.Val, 0) && visit(node.Right, node.Val, 1)
}*/

var printf = fmt.Printf

func isValidBST(node *TreeNode) bool {
	isValid, _, _ := visit(node)
	return isValid
}

func visit(node *TreeNode) (bool, int, int) {
	if node.Left == nil && node.Right == nil {
		// debug
	//	printf("leaf node, val: %d\n", node.Val)
		return true, node.Val, node.Val
	}

	if node.Right == nil {
		isValid, smallest, biggest := visit(node.Left)
		//printf("node val: %d, left is valid? %b, greatest:%d\n", node.Val, isValid, greatest)
		if !isValid || biggest >= node.Val {
			return false, 0,0
		}
		return true, smallest, node.Val
	}

	if node.Left == nil {
		isValid, smallest, biggest := visit(node.Right)
	//	printf("node val: %d, right is valid? %b, smallest:%d\n", node.Val, isValid, smallest)
		if !isValid || smallest <= node.Val {
			return false, 0,0
		}
		return true, node.Val, biggest
	}

	isValidL, smallestL, biggestL := visit(node.Left)
	isValidR, smallestR, biggestR := visit(node.Right)

	if isValidL && isValidR && biggestL < node.Val && smallestR > node.Val {
		//printf("val: %d, biggest: %d, smallest: %d\n", node.Val, biggest, smallest)
		return true, smallestL, biggestR 
	}
	return false, 0,0
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
