/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var dfs func(*TreeNode)
	dfs = func(t *TreeNode) {
		if t == nil {
			return
		}
		dfs(t.Right)
		sum += t.Val
		t.Val = sum
		dfs(t.Left)
	}
	dfs(root)
	return root
}

// @lc code=end

