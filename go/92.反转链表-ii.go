/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	//虚拟头结点
	prev := &ListNode{Next: head}
	prevHead := prev
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	tmpHead := prev.Next
	//本质上是固定住prev
	for i := 0; i < right-left; i++ {
		next := tmpHead.Next
		tmpHead.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return prevHead.Next
}

// @lc code=end

