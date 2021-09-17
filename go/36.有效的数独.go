/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var row, column [9][9]int
	var total [3][3][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			info := board[i][j] - '1'
			//更新行
			row[i][info]++
			//更新列
			column[j][info]++
			//更新小方格
			total[i/3][j/3][info]++

			if row[i][info] > 1 || column[j][info] > 1 || total[i/3][j/3][info] > 1 {
				return false
			}

		}
	}
	return true
}

// @lc code=end

