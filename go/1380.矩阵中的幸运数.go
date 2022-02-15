/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	ans := make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	rows := make([]int, m)
	columns := make([]int, n)
	for i := 0; i < m; i++ {
		rows[i] = math.MaxInt32
		for j := 0; j < n; j++ {
			rows[i] = min(rows[i], matrix[i][j])
			columns[j] = max(columns[j], matrix[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == rows[i] && matrix[i][j] == columns[j] {
				ans = append(ans, matrix[i][j])
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

