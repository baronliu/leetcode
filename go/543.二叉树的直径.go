/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 *
 * https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (45.04%)
 * Likes:    270
 * Dislikes: 0
 * Total Accepted:    35.8K
 * Total Submissions: 73.8K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。
 *
 * 示例 :
 * 给定二叉树
 *
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 *
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
 *
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
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	deep(root, &result)
	return result
}

func deep(node *TreeNode, result *int) int {
	if node == nil {
		return 0
	}

	var deepLeft, deepRight int

	if node.Left != nil {
		deepLeft = deep(node.Left, result) + 1
	}

	if node.Right != nil {
		deepRight = deep(node.Right, result) + 1
	}

	if deepLeft+deepRight > *result {
		*result = deepLeft + deepRight
	}

	return max(deepLeft, deepRight)
}

func max(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}

// @lc code=end

