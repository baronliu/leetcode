/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	ans := 0

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}

			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}

	return ans * ans
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}

	if a > c {
		a = c
	}
	return a
}

// @lc code=end

