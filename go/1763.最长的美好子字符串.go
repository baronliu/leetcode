/*
 * @lc app=leetcode.cn id=1763 lang=golang
 *
 * [1763] 最长的美好子字符串
 */

// @lc code=start
func longestNiceSubstring(s string) string {
	ans := ""
	length := len(s)
	if length == 1 {
		return ans
	}
	memo := make([][]int, length+1)
	for i := 0; i < length+1; i++ {
		memo[i] = make([]int, 52)
	}
	for i := 0; i < length; i++ {
		copy(memo[i+1], memo[i])
		memo[i+1][convert(s[i])]++
	}

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if check(memo[i], memo[j+1]) {
				if j+1-i > len(ans) {
					ans = s[i : j+1]
				}
			}
		}
	}

	return ans
}

func convert(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b - 'a')
	}

	if b >= 'A' && b <= 'Z' {
		return int(b - 'A' + 26)
	}
	return 0
}

func check(a, b []int) bool {
	for i := 0; i < 26; i++ {
		x := b[i] - a[i]
		y := b[i+26] - a[i+26]
		if (x == 0 && y > 0) || (x > 0 && y == 0) {
			return false
		}
	}
	return true
}

// @lc code=end

