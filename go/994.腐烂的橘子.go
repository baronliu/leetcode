/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 *
 * https://leetcode-cn.com/problems/rotting-oranges/description/
 *
 * algorithms
 * Easy (45.30%)
 * Likes:    150
 * Dislikes: 0
 * Total Accepted:    18K
 * Total Submissions: 35.8K
 * Testcase Example:  '[[2,1,1],[1,1,0],[0,1,1]]'
 *
 * 在给定的网格中，每个单元格可以有以下三个值之一：
 *
 *
 * 值 0 代表空单元格；
 * 值 1 代表新鲜橘子；
 * 值 2 代表腐烂的橘子。
 *
 *
 * 每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
 *
 * 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[[2,1,1],[1,1,0],[0,1,1]]
 * 输出：4
 *
 *
 * 示例 2：
 *
 * 输入：[[2,1,1],[0,1,1],[1,0,1]]
 * 输出：-1
 * 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
 *
 *
 * 示例 3：
 *
 * 输入：[[0,2]]
 * 输出：0
 * 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= grid.length <= 10
 * 1 <= grid[0].length <= 10
 * grid[i][j] 仅为 0、1 或 2
 *
 *
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	result := 0
	if len(grid) < 1 {
		return result
	}
	var queue [][]int

	//1.初始化腐烂的橘子
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				point := []int{i, j, 0}
				queue = append(queue, point)
			}
		}
	}

	//2.循环腐烂
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		result = point[2]
		//上
		if point[0] >= 1 && grid[point[0]-1][point[1]] == 1 {
			grid[point[0]-1][point[1]] = 2
			queue = append(queue, []int{point[0] - 1, point[1], point[2] + 1})
		}
		//下
		if point[0] < len(grid)-1 && grid[point[0]+1][point[1]] == 1 {
			grid[point[0]+1][point[1]] = 2
			queue = append(queue, []int{point[0] + 1, point[1], point[2] + 1})
		}
		//左
		if point[1] >= 1 && grid[point[0]][point[1]-1] == 1 {
			grid[point[0]][point[1]-1] = 2
			queue = append(queue, []int{point[0], point[1] - 1, point[2] + 1})
		}
		//右
		if point[1] < len(grid[0])-1 && grid[point[0]][point[1]+1] == 1 {
			grid[point[0]][point[1]+1] = 2
			queue = append(queue, []int{point[0], point[1] + 1, point[2] + 1})
		}
	}

	//3.检查新鲜的
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return result
}

// @lc code=end

