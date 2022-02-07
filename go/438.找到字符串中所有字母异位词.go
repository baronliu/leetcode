/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	if len(s) < len(p) {
		return ans
	}
	var window, base [26]int
	for i := 0; i < len(p); i++ {
		base[int(p[i]-'a')]++
		window[int(s[i]-'a')]++
	}
	if base == window {
		ans = append(ans, 0)
	}

	for i := 1; i <= len(s)-len(p); i++ {
		window[int(s[i-1]-'a')]--
		window[int(s[i+len(p)-1]-'a')]++
		if window == base {
			ans = append(ans, i)
		}
	}

	return ans
}

// @lc code=end

