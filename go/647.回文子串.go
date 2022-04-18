/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	ans := 1
	for i := 0; i < length-1; i++ {
		ans += count(s, i, i) + count(s, i, i+1)
	}

	return ans
}

func count(s string, start, end int) int {
	ans := 0
	for start >= 0 && end < len(s) && s[start] == s[end] {
		ans++
		start--
		end++
	}
	return ans
}

// @lc code=end

