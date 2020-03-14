/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (42.42%)
 * Likes:    487
 * Dislikes: 0
 * Total Accepted:    61.7K
 * Total Submissions: 140.1K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */

// @lc code=start
//方法1：动态规划
func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	dp := make([]int, len(nums))
	for k := range dp {
		dp[k] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}

//方法2：贪心思想 + 二分查找
// func lengthOfLIS(nums []int) int {
// 	l := len(nums)
// 	if l <= 1 {
// 		return l
// 	}
// 	var dp []int
// 	dp = append(dp, nums[0])
// 	for i := 1; i < len(nums); i++ {
// 		lastOne := dp[len(dp)-1]
// 		if nums[i] > lastOne {
// 			dp = append(dp, nums[i])
// 		} else {
// 			start, end := 0, len(dp)-1
// 			for start < end {
// 				mid := (start + end) / 2
// 				if dp[mid] >= nums[i] {
// 					end = mid
// 				} else {
// 					start = mid + 1
// 				}
// 			}
// 			dp[start] = nums[i]
// 		}
// 	}
// 	return len(dp)
// }

func max(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}

// @lc code=end

