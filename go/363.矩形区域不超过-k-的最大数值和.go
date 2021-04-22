/*
 * @lc app=leetcode.cn id=363 lang=golang
 *
 * [363] 矩形区域不超过 K 的最大数值和
 */

// @lc code=start
func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	ans := -math.MaxInt32
	for a := 1; a <= m; a++ {
		for b := 1; b <= n; b++ {
			for c := a; c <= m; c++ {
				for d := b; d <= n; d++ {
					cur := dp[c][d] - dp[a-1][d] - dp[c][b-1] + dp[a-1][b-1]
					if cur <= k {
						ans = max(ans, cur)
					}
				}
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

