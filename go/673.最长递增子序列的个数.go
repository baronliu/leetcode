/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}

	ans, maxLength := 1, 1
	dp, cnt := make([]int, length), make([]int, length)
	dp[0], cnt[0] = 1, 1
	for i := 1; i < length; i++ {
		dp[i], cnt[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				cnt[i] = cnt[j]
			} else if dp[j]+1 == dp[i] {
				cnt[i] += cnt[j]
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxLength {
			ans += cnt[i]
		}
	}
	return ans
}

// @lc code=end

