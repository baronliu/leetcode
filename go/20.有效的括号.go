/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (38.81%)
 * Likes:    1421
 * Dislikes: 0
 * Total Accepted:    219K
 * Total Submissions: 533.5K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

// @lc code=start

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	var stack []string

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == "(" || c == "[" || c == "{" {
			//入栈
			stack = append(stack, c)
		} else if len(stack) < 1 {
			return false
		} else {
			info := stack[len(stack)-1]
			if (info == "(" && c != ")") || (info == "[" && c != "]") || (info == "{" && c != "}") {
				return false
			}
			stack = stack[:(len(stack) - 1)]
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end

