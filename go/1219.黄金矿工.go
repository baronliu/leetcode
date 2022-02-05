/*
 * @lc app=leetcode.cn id=1219 lang=golang
 *
 * [1219] 黄金矿工
 */

// @lc code=start
func getMaximumGold(grid [][]int) int {
	ans := 0
	var dfs func(int, int, int)
	dfs = func(i, j, cur int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
			return
		}
		v := grid[i][j]
		cur += v
		ans = max(ans, cur)
		grid[i][j] = 0
		dfs(i-1, j, cur)
		dfs(i+1, j, cur)
		dfs(i, j-1, cur)
		dfs(i, j+1, cur)
		grid[i][j] = v
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dfs(i, j, 0)
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

