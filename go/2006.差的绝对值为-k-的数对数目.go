/*
 * @lc app=leetcode.cn id=2006 lang=golang
 *
 * [2006] 差的绝对值为 K 的数对数目
 */

// @lc code=start
func countKDifference(nums []int, k int) int {
	length := len(nums)
	ans := 0
	if length == 1 {
		return ans
	}

	memo := make(map[int]int)
	for i := 0; i < length; i++ {
		ans += memo[nums[i]+k] + memo[nums[i]-k]
		memo[nums[i]]++
	}

	return ans
}

// @lc code=end

