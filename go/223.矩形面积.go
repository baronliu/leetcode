/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	m := (ax2 - ax1) * (ay2 - ay1)
	n := (bx2 - bx1) * (by2 - by1)
	x := min(ax2, bx2) - max(ax1, bx1)
	y := min(ay2, by2) - max(ay1, by1)

	return m + n - max(x, 0)*max(y, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

