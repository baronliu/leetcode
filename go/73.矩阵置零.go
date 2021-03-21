/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	mArray, nArray := make([]int, 0), make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				//记录下目标行和列
				mArray = append(mArray, i)
				nArray = append(nArray, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if inArray(mArray, i) || inArray(nArray, j) {
				matrix[i][j] = 0
			}
		}
	}
}

func inArray(arr []int, target int) bool {
	for _, a := range arr {
		if a == target {
			return true
		}
	}
	return false
}

// @lc code=end

