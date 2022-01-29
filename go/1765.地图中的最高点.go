/*
 * @lc app=leetcode.cn id=1765 lang=golang
 *
 * [1765] 地图中的最高点
 */

// @lc code=start
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ans := make([][]int, m)
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				queue = append(queue, []int{i, j})
			} else {
				ans[i][j] = -1
			}
		}
	}
	h := 1
	dict := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		newQueue := make([][]int, 0)
		for i := 0; i < len(queue); i++ {
			v := queue[i]
			for _, d := range dict {
				newM, newN := v[0]+d[0], v[1]+d[1]
				if newM < 0 || newN < 0 || newM >= m || newN >= n || ans[newM][newN] != -1 {
					continue
				}
				ans[newM][newN] = h
				newQueue = append(newQueue, []int{newM, newN})
			}
		}
		h++
		queue = newQueue
	}

	return ans
}

// @lc code=end

