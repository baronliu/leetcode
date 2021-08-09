/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	nums := []int{2, 3, 5}
	for _, num := range nums {
		for n%num == 0 {
			n /= num
		}
	}
	return n == 1
}

// @lc code=end

