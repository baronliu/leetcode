/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) < 1 {
		return 0
	}
	var retNum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				retNum++
				//把所有相关的1全部置为0
				dfs(grid, i, j)
			}
		}
	}

	return retNum
}

func dfs(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	//四个方向遍历
	//上
	if i > 0 {
		dfs(grid, i-1, j)
	}
	//下
	if i < len(grid)-1 {
		dfs(grid, i+1, j)
	}
	//左
	if j > 0 {
		dfs(grid, i, j-1)
	}
	//右
	if j < len(grid[0])-1 {
		dfs(grid, i, j+1)
	}
}

// @lc code=end

