/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (59.51%)
 * Likes:    101
 * Dislikes: 0
 * Total Accepted:    10.1K
 * Total Submissions: 16.5K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 *
 * 示例:
 *
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 *
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var result []int
	return appendResult(root, result, 0)
}

func appendResult(t *TreeNode, r []int, level int) []int {
	if t == nil {
		return r
	}

	if len(r) <= level {
		r = append(r, t.Val)
	}

	r = appendResult(t.Right, r, level+1)
	r = appendResult(t.Left, r, level+1)

	return r
}

