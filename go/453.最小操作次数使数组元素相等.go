/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	minInt := math.MaxInt32
	for _, num := range nums {
		minInt = min(minInt, num)
	}
	ans := 0
	for _, num := range nums {
		ans += num - minInt
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

