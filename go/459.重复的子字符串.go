/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	return kmp(s+s, s)
}

func kmp(query, pattern string) bool {
	n, m := len(query), len(pattern)
	fail := make([]int, m)
	for i := 0; i < m; i++ {
		fail[i] = -1
	}
	for i := 1; i < m; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	match := -1
	for i := 1; i < n-1; i++ {
		for match != -1 && pattern[match+1] != query[i] {
			match = fail[match]
		}
		if pattern[match+1] == query[i] {
			match++
			if match == m-1 {
				return true
			}
		}
	}
	return false
}

// @lc code=end

