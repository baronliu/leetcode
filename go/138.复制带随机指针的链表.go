/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

/**
 * 思路：先不管随机指针的属性，就当是复制一个普通的链表，那么问题核心就是随机指针的复制
 * 对于随机指针的复制，核心问题点是我们不能提前知道目标地址，那么理所当然考虑两次遍历
 * 第一次遍历的时候记录下每个节点的新地址和旧地址的隐射关系，第二次遍历把旧的随机节点替换成新的随机节点即可
 */
func copyRandomList(head *Node) *Node {
	cur := head
	res := new(Node)
	p := res
	dict := make(map[*Node]*Node)
	for cur != nil {
		v := &Node{
			Val: cur.Val,
		}
		dict[cur] = v
		p.Next = v
		p = p.Next
		cur = cur.Next
	}

	cur = head
	p = res.Next
	for cur != nil {
		p.Random = dict[cur.Random]
		p = p.Next
		cur = cur.Next
	}

	return res.Next
}

// @lc code=end

