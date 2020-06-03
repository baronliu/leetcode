/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	var queue []*TreeNode
	queue = append(queue, root)
	index := 0
	for len(queue) > 0 {
		l := len(queue)
		//一定有这么多个元素
		temp := make([]int, l)
		var newQueue []*TreeNode
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
			if index%2 != 0 {
				//逆序
				temp[l-1-i] = queue[i].Val
			} else {
				//顺序
				temp[i] = queue[i].Val
			}
		}
		index++
		result = append(result, temp)
		queue = newQueue
	}

	return result
}

// @lc code=end

