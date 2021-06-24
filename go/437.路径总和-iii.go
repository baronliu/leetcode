/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	//1.存放记录一条路径上的所有前缀和出现的次数
	dict := make(map[int]int)
	dict[0] = 1
	var ans int

	var dfs func(*TreeNode, int)

	dfs = func(n *TreeNode, pre int) {
		if n == nil {
			return
		}
		pre = pre + n.Val
		ans += dict[pre-targetSum]
		//加数据应该放在后面，避免target=0的情况，自己-自己
		dict[pre]++
		dfs(n.Left, pre)
		dfs(n.Right, pre)
		dict[pre]--
	}

	dfs(root, 0)
	return ans
}

// @lc code=end

