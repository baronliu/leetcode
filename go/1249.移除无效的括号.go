/*
 * @lc app=leetcode.cn id=1249 lang=golang
 *
 * [1249] 移除无效的括号
 */

// @lc code=start
func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0)
	b := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				b[i] = byte(' ')
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	for i := 0; i < len(stack); i++ {
		b[stack[i]] = byte(' ')
	}

	return strings.ReplaceAll(string(b), " ", "")
}

// @lc code=end

