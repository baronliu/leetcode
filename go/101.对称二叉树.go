/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	//1.递归法
	// return check(root, root)

	//2.迭代法
	var queue []*TreeNode
	if root == nil {
		return true
	}
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		n1, n2 := queue[0], queue[1]
		//弹出最前面两个元素
		queue = queue[2:]
		if n1 == nil && n2 == nil {
			continue
		}

		if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}
		queue = append(queue, n1.Left, n2.Right)
		queue = append(queue, n1.Right, n2.Left)
	}

	return true
}

func check(a, b *TreeNode) bool {
	//左右都为空，表示对称
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && check(a.Left, b.Right) && check(a.Right, b.Left)
}

// @lc code=end

