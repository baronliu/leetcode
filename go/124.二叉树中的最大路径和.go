/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64
	maxValue(root, &ans)
	return ans
}

func maxValue(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	//因为是任意节点，所以不一定是叶子节点
	left := max(maxValue(root.Left, ans), 0)
	right := max(maxValue(root.Right, ans), 0)

	*ans = max(*ans, left+right+root.Val)

	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

