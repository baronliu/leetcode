/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) < 2 {
		return nums[0]
	}
	return max(robOne(nums[:len(nums)-1]), robOne(nums[1:]))
}

func robOne(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]

		if i > 1 {
			cur = dp[i-2] + nums[i]
		}

		if i > 2 {
			cur = max(dp[i-3]+nums[i], cur)
		}
		dp[i] = cur
		ret = max(ret, cur)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

