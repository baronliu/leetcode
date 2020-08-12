/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var queue, newQueue []*Node

	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue = nil
		for i := 0; i < len(queue); i++ {
			if i != len(queue)-1 {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}

		queue = newQueue
	}
	return root
}

// @lc code=end

