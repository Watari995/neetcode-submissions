/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) int 
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left == -1 || right == -1 {
			return -1
		}
		// 各nodeの高さが合ってなければ-1を返す
		if left - right > 1 || right - left > 1 {
			return -1
		}
		return max(left, right) + 1
	}   
	return dfs(root) != -1
}
