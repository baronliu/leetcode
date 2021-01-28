/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	dp := []int{1, 1}

	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}

	return dp[n]
}

// @lc code=end

