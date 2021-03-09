/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(S string) string {
	//1.长度小于等于1时直接返回当前字符串
	if len(S) <= 1 {
		return S
	}
	//2.模拟栈
	stack := []byte{S[0]}
	for i := 1; i < len(S); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}

// @lc code=end

