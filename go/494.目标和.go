/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	var result int
	if len(nums) < 1 {
		return result
	}
	find(nums, 0, 0, S, &result)

	return result
}

func find (nums []int, current, t, s int, result *int) {
	//结束
	if (current >= len(nums)) {
		if t == s {
			*result = *result + 1
		}
		return
	}

	//+
	find(nums, current+1, t + nums[current], s, result)

	//-
	find(nums, current+1, t - nums[current], s, result)
}



// @lc code=end

