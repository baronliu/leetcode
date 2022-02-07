/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 1
	}
	dp := make([]int, length)
	dp[0] = 1
	ans := 1
	for i := 1; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		ans = max(ans, dp[i])
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

