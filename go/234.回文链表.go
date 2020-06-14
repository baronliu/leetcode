/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (41.75%)
 * Likes:    523
 * Dislikes: 0
 * Total Accepted:    95.2K
 * Total Submissions: 226K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 * 
 * 示例 1:
 * 
 * 输入: 1->2
 * 输出: false
 * 
 * 示例 2:
 * 
 * 输入: 1->2->2->1
 * 输出: true
 * 
 * 
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    //获取值数组
    var values []int

    for head != nil {
        values = append(values, head.Val)
        head = head.Next
    }

    //双指针
    i, j := 0, len(values) -1
    for i<=j {
        if values[i] != values [j] {
            return false
        }
        i += 1
        j -= 1
    }

    return true
}
// @lc code=end

