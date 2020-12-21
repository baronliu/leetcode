/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}
	//初始化dp数组
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var cur int
			if i > 0 && j > 0 {
				cur = min(dp[i-1][j], dp[i][j-1])
			} else if i > 0 {
				cur = dp[i-1][j]
			} else if j > 0 {
				cur = dp[i][j-1]
			} else {
				cur = 0
			}

			dp[i][j] = grid[i][j] + cur
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

