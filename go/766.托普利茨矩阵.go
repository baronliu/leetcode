/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])

	for i := m - 1; i >= 0; i-- {
		a, b := i, 0
		for a <= m-1 && b <= n-1 {
			if matrix[a][b] != matrix[i][0] {
				return false
			}
			a++
			b++
		}
	}

	for j := 1; j <= n-1; j++ {
		a, b := 0, j
		for a <= m-1 && b <= n-1 {
			if matrix[a][b] != matrix[0][j] {
				return false
			}
			a++
			b++
		}
	}

	return true
}

// @lc code=end

