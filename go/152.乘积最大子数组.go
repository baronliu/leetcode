/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (38.16%)
 * Likes:    571
 * Dislikes: 0
 * Total Accepted:    67.1K
 * Total Submissions: 170.5K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */

// @lc code=start

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	maxResult := nums[0]
	iMax := nums[0]
	iMin := nums[0]

	for i := 1; i < len(nums); i++ {
		//交换最大值和最小值
		if nums[i] < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(iMax*nums[i], nums[i])
		iMin = min(iMin*nums[i], nums[i])

		//求解最大值
		maxResult = max(maxResult, iMax)
	}
	return maxResult
}

// @lc code=end

