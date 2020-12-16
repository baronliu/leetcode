/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m < 1 {
		return nil
	}
	n := len(matrix[0])
	if n < 1 {
		return nil
	}
	//总数
	total := m * n
	ret := make([]int, 0)
	i := 0
	for total > 0 {
		left, right, top, bottom := i, n-i-1, i, m-i-1

		for j := left; j <= right; j++ {
			ret = append(ret, matrix[top][j])
			total--
		}

		for j := top + 1; j <= bottom; j++ {
			ret = append(ret, matrix[j][right])
			total--
		}

		for j := right - 1; j >= left && bottom > top; j-- {
			ret = append(ret, matrix[bottom][j])
			total--
		}

		for j := bottom - 1; j > top && right > left; j-- {
			ret = append(ret, matrix[j][left])
			total--
		}

		i++
	}
	return ret
}

// @lc code=end

