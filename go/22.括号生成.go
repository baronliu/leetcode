/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (70.99%)
 * Likes:    973
 * Dislikes: 0
 * Total Accepted:    118.3K
 * Total Submissions: 157.2K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例：
 *
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var result []string
	if n == 0 {
		return result
	}
	dfs(n, n, "", &result)
	return result
}

func dfs(left int, right int, current string, result *[]string) {
	if left == 0 && right == 0 {
		//到底后写入
		*result = append(*result, current)
	}

	//左括号用完了
	if left > 0 {
		dfs(left-1, right, current+"(", result)
	}

	//右括号
	if right > left {
		dfs(left, right-1, current+")", result)
	}
}

// @lc code=end

