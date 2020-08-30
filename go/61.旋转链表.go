/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * 思路：先把全部的节点存下来，然后根据要移动的数量，计算出来头节点，从头结点拼到slice的结尾，再接上slice开头
 */
func rotateRight(head *ListNode, k int) *ListNode {
	nodes := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	if len(nodes) <= 1 {
		return head
	}
	k = k % len(nodes)
	if k == 0 {
		return head
	}
	for i := len(nodes) - k; i < len(nodes); i++ {
		if i == len(nodes)-1 {
			nodes[i].Next = nodes[0]
		} else {
			nodes[i].Next = nodes[i+1]
		}
	}

	for i := 0; i < len(nodes)-k; i++ {
		if i == len(nodes)-k-1 {
			nodes[i].Next = nil
		} else {
			nodes[i].Next = nodes[i+1]
		}
	}
	return nodes[len(nodes)-k]
}

// @lc code=end

