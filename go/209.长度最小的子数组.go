/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	i, sum, result := 0, 0, len(nums) + 1

	for j:=0; j<len(nums); j++ {
		sum += nums[j]
		for sum >= s {
			if j - i + 1 < result {
				result = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}

	if result == len(nums) + 1 {
		return 0
	}
	return result
}
// @lc code=end

