/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)

		if left == -1 || right == -1 {
			return -1
		}

		if abs(left-right) > 1 {
			return -1
		}

		return max(left, right) + 1 
	}

	return depth(root) != -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
