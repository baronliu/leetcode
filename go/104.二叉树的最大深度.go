/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var result int

    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    if left >= right {
        result = left
    } else {
        result = right
    }

    return result + 1
}
// @lc code=end

