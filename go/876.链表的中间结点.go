/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var s []*ListNode
	s = append(s, head)
	for head.Next != nil {
		s = append(s, head.Next)
		head = head.Next
	}
	index := len(s) / 2
	return s[index]
}

