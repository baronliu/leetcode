/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	//1.初始化dp数组
	dp := make([][][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 1; i <= length; i++ {
		zeros := strings.Count(strs[i-1], "0")
		ones := len(strs[i-1]) - zeros
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if j >= zeros && k >= ones {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-zeros][k-ones]+1)
				}
			}
		}
	}
	return dp[length][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

