/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var find func(int, int, int) bool

	find = func(l, a, b int) bool {
		if l == len(word) {
			return true
		}

		if a < 0 || b < 0 || a >= m || b >= n {
			return false
		}

		if used[a][b] || word[l] != board[a][b] {
			return false
		}

		used[a][b] = true

		if find(l+1, a, b+1) || find(l+1, a, b-1) || find(l+1, a+1, b) || find(l+1, a-1, b) {
			return true
		}

		used[a][b] = false

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if find(0, i, j) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

