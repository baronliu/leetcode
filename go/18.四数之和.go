/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (34.17%)
 * Total Accepted:    16.4K
 * Total Submissions: 46.9K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
// package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	l := len(nums)

	for m := 0; m < l-3; m++ {

		if m > 0 && nums[m] == nums[m-1] {
			continue
		}

		for n := m + 1; n < l-2; n++ {
			if n > m+1 && nums[n] == nums[n-1] {
				continue
			}
			i, j := n+1, l-1
			for i < j {
				if nums[m]+nums[n]+nums[i]+nums[j] == target {
					result = append(result, []int{nums[m], nums[n], nums[i], nums[j]})
					i++
					j--
					for i < j && nums[i] == nums[i-1] {
						i++
					}

					for i < j && nums[j] == nums[j+1] {
						j--
					}

				} else if nums[m]+nums[n]+nums[i]+nums[j] > target {
					j--
				} else {
					i++
				}
			}
		}
	}

	return result
}
