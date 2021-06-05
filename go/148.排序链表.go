/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	r := sortList(slow.Next)
	slow.Next = nil
	return merge(sortList(head), r)
}

func merge(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	head := &ListNode{}
	cur := head

	for a != nil && b != nil {
		if a.Val <= b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a != nil {
		cur.Next = a
	}

	if b != nil {
		cur.Next = b
	}
	return head.Next
}

// @lc code=end

