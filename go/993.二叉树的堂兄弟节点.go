/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
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
func isCousins(root *TreeNode, x int, y int) bool {
	foundX, foundY := false, false
	deepX, deepY := 0, 0
	parentX, parentY := 0, 0

	var deep func(*TreeNode, int, int)
	deep = func(n *TreeNode, p, d int) {
		if n == nil {
			return
		}

		if n.Val == x {
			deepX = d
			foundX = true
			parentX = p
		} else if n.Val == y {
			deepY = d
			foundY = true
			parentY = p
		}

		if foundX && foundY {
			return
		}

		deep(n.Left, n.Val, d+1)
		deep(n.Right, n.Val, d+1)
	}

	deep(root, 0, 0)

	return deepX == deepY && parentX != parentY
}

// @lc code=end

