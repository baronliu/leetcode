/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var lastVisit, currentVisit *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		currentVisit = stack[len(stack)-1]
		//没有右侧节点 或 右侧节点已经被遍历过
		if currentVisit.Right == nil || lastVisit == currentVisit.Right {
			result = append(result, currentVisit.Val)
			stack = stack[:len(stack)-1]
			lastVisit = currentVisit
		} else {
			root = currentVisit.Right
		}
	}
	return result
}

// @lc code=end

