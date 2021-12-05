/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	ans := 0
	dp := make(map[int]int)
	for i := 1; i <= n; i++ {
		x := i
		if x%5 == 0 {
			dp[x] = dp[x/5] + 1
			ans += dp[x]
		}
	}
	return ans
}

// @lc code=end

