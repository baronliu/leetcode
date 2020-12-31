/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	return merge(mergeKLists(lists[:(length/2)]), mergeKLists(lists[(length/2):]))

	// ans := lists[0]
	// for i:=1; i<len(lists); i++ {
	//     ans = merge(ans, lists[i])
	// }
	// return ans
}

func merge(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	head := &ListNode{}
	prev := head
	for a != nil && b != nil {
		var cur *ListNode
		if a.Val < b.Val {
			cur = &ListNode{Val: a.Val}
			a = a.Next
		} else {
			cur = &ListNode{Val: b.Val}
			b = b.Next
		}
		prev.Next = cur
		prev = cur
	}

	if a == nil {
		prev.Next = b
	}

	if b == nil {
		prev.Next = a
	}

	return head.Next
}

// @lc code=end

