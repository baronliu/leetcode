/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除与获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	newNums := make([]int, 10001)
	dp := make([]int, 10001)
	for _, num := range nums {
		newNums[num] = newNums[num] + num
	}
	dp[1] = newNums[1]

	for i := 2; i <= 10000; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+newNums[i])
	}

	return dp[10000]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

