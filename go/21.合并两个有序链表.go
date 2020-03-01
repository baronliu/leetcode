/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (52.89%)
 * Total Accepted:    47.3K
 * Total Submissions: 89.4K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// package leetcode

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var p1, p2, current *ListNode
	head := new(ListNode)
	p1, p2, current = l1, l2, head
	for p1 != nil || p2 != nil {
		//如果一方指针遍历完全，另一方直接接到底部
		if p1 == nil {
			current.Next = p2
			break
		}

		if p2 == nil {
			current.Next = p1
			break
		}

		if p1.Val > p2.Val {
			current.Next = p2
			p2 = p2.Next
		} else {
			current.Next = p1
			p1 = p1.Next
		}
		current = current.Next

	}
	return head.Next
}
