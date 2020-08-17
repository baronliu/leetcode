ß/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    //只有 0 或 1 个元素
    for head == nil || head.Next == nil {
        return head
    }
    a := head
    b := head.Next
    c := head.Next
    cur := head.Next.Next
    i := 1
    for cur != nil {
        if i%2 != 0 {
            a.Next = cur
            a = cur
        } else {
            b.Next = cur
            b = cur
        }
        cur = cur.Next
        i++
    }
    a.Next = c
    b.Next = nil
    return head
}
// @lc code=end

