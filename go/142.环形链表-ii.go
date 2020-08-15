/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
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
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	a, b := head, head
	for a != nil && b != nil && a.Next != nil {
		a = a.Next.Next
		b = b.Next
		if a == b {
			b = head
			for a != b {
				a = a.Next
				b = b.Next
			}
			return a
		}
	}
	return nil
}

// @lc code=end

