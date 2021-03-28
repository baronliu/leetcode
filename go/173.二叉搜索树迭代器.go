/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
type BSTIterator struct {
	Cur int   //当前索引到第几个元素
	Arr []int //元素数组化
}

func Constructor(root *TreeNode) BSTIterator {
	item := BSTIterator{}
	item.Cur = -1
	item.Arr = inorder(root)
	return item
}

func inorder(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	left := inorder(root.Left)
	right := inorder(root.Right)

	return append(append(left, root.Val), right...)
}

func (this *BSTIterator) Next() int {
	this.Cur += 1
	return this.Arr[this.Cur]
}

func (this *BSTIterator) HasNext() bool {
	return this.Cur < len(this.Arr)-1
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

