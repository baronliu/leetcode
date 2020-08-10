/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	//叶子节点
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	left, right := false, false
	if root.Left != nil {
		left = hasPathSum(root.Left, sum-root.Val)
	}

	if root.Right != nil {
		right = hasPathSum(root.Right, sum-root.Val)
	}

	return left || right
}

// @lc code=end

