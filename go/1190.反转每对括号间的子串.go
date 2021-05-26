/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 */

// @lc code=start
func reverseParentheses(s string) string {
	if len(s) <= 1 {
		return s
	}
	tmp := ""
	stack := make([]string, 0)
	for _, b := range s {
		if b == '(' {
			stack = append(stack, tmp)
			tmp = ""
		} else if b == ')' {
			tmp = stack[len(stack)-1] + convert(tmp)
			stack = stack[:len(stack)-1]
		} else {
			tmp += string(b)
		}
	}
	return tmp
}

func convert(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// @lc code=end

