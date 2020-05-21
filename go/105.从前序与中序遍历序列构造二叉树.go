/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (65.47%)
 * Likes:    447
 * Dislikes: 0
 * Total Accepted:    67.2K
 * Total Submissions: 102.3K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//取当前根节点
	root := preorder[0]

	//找到分割 左子树 和 右子树 的位置
	splitIndex := 0
	for k, v := range inorder {
		if v == root {
			splitIndex = k
		}
	}

	t := &TreeNode{}
	t.Val = root
	//递归构建 左右 子树
	t.Left = buildTree(preorder[1:1+splitIndex], inorder[:splitIndex])
	t.Right = buildTree(preorder[1+splitIndex:], inorder[1+splitIndex:])
	return t
}

// @lc code=end

