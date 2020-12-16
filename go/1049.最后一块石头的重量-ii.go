/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
//背包问题
func lastStoneWeightII(stones []int) int {
	length := len(stones)
	if length < 2 {
		return stones[0]
	}
	total := 0
	for i := 0; i < length; i++ {
		total += stones[i]
	}
	sum := total / 2
	dp := make([]int, sum+1)
	for i := 0; i < length; i++ {
		for j := sum; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return total - 2*dp[sum]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

