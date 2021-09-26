/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start`
func minDistance(word1 string, word2 string) int {
	maxCount := 0
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		maxCount = 0
	} else {
		dp := make([][]int, m+1)
		for i := 0; i <= m; i++ {
			dp[i] = make([]int, n+1)
		}

		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i][j-1], dp[i-1][j])
				}
			}
		}
		maxCount = dp[m][n]
	}

	return m + n - 2*maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

