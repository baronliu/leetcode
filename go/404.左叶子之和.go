/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans, left, right int
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			left = root.Left.Val
		} else {
			left = sumOfLeftLeaves(root.Left)
		}
	}

	right = sumOfLeftLeaves(root.Right)
	return left + right
}

// @lc code=end

