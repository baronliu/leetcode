/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}

	//总和不为偶数
	if sum%2 != 0 {
		return false
	}

	dp := make([][]int, length+1)

	for i := 0; i < length+1; i++ {
		dp[i] = make([]int, sum/2+1)
	}
	for i := 1; i <= length; i++ {
		for j := 1; j <= sum/2; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i-1]]+nums[i-1])
			}
		}
	}

	return dp[length][sum/2] == sum/2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

