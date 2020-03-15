/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 *
 * https://leetcode-cn.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (55.24%)
 * Likes:    208
 * Dislikes: 0
 * Total Accepted:    28.1K
 * Total Submissions: 45.4K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地)
 * 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
 *
 * 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
 *
 * 示例 1:
 *
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 *
 * 对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
 *
 * 示例 2:
 *
 *
 * [[0,0,0,0,0,0,0,0]]
 *
 * 对于上面这个给定的矩阵, 返回 0。
 *
 * 注意: 给定的矩阵grid 的长度和宽度都不超过 50。
 *
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	maxCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count := getCount(grid, i, j)
				if count > maxCount {
					maxCount = count
				}
			}
		}
	}
	return maxCount
}

//1. BFS
func getCount(grid [][]int, i int, j int) int {
	//坐标序列
	var queue [][]int
	result := 0
	queue = append(queue, []int{i, j})
	for len(queue) > 0 {
		item := queue[0]
		m, n := item[0], item[1]
		queue = queue[1:]
		if grid[m][n] == 1 {
			grid[m][n] = 2
			result++
			//添加新四个方向上的新元素
			//上
			if m >= 1 && grid[m-1][n] == 1 {
				queue = append(queue, []int{m - 1, n})
			}
			//下
			if m < len(grid)-1 && grid[m+1][n] == 1 {
				queue = append(queue, []int{m + 1, n})
			}
			//左
			if n >= 1 && grid[m][n-1] == 1 {
				queue = append(queue, []int{m, n - 1})
			}
			//右
			if n < len(grid[0])-1 && grid[m][n+1] == 1 {
				queue = append(queue, []int{m, n + 1})
			}
		}
	}
	return result
}

//2. DFS
// func getCount(grid [][]int, m int, n int) int {
// 	result := 0
// 	if grid[m][n] == 1 {
// 		grid[m][n] = 2
// 		result++
// 		//添加新四个方向上的新元素
// 		//上
// 		if m >= 1 && grid[m-1][n] == 1 {
// 			result += getCount(grid, m-1, n)
// 		}
// 		//下
// 		if m < len(grid)-1 && grid[m+1][n] == 1 {
// 			result += getCount(grid, m+1, n)
// 		}
// 		//左
// 		if n >= 1 && grid[m][n-1] == 1 {
// 			result += getCount(grid, m, n-1)
// 		}
// 		//右
// 		if n < len(grid[0])-1 && grid[m][n+1] == 1 {
// 			result += getCount(grid, m, n+1)
// 		}
// 	}
// 	return result
// }

// @lc code=end

