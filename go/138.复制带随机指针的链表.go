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

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dict := make(map[*Node]*Node)

	temp := head
	prev := &Node{}
	retNode := prev
	for temp != nil {
		n := getCloneNode(dict, temp)
		n.Random = getCloneNode(dict, temp.Random)
		prev.Next = n
		prev = n
		temp = temp.Next
	}

	return retNode.Next
}

func getCloneNode(dict map[*Node]*Node, node *Node) *Node {
	if newNode, ok := dict[node]; ok {
		return newNode
	} else {
		var newNode *Node
		if node != nil {
			newNode = &Node{Val: node.Val}
		}
		dict[node] = newNode
		return newNode
	}
}

// @lc code=end

