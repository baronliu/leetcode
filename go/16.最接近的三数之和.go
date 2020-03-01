import "sort"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (39.06%)
 * Total Accepted:    21.9K
 * Total Submissions: 55.1K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */

func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	sort.Ints(nums)
	min := nums[0] + nums[1] + nums[2]
	max := nums[l-1] + nums[l-2] + nums[l-3]
	result, minR := min, max-min
	if target <= min {
		return min
	} else if target >= max {
		return max
	} else {
		for i := 0; i < l-2; i++ {
			j, k, current := i+1, l-1, 0 //两个指针去遍历
			for j < k {
				current = nums[i] + nums[j] + nums[k]
				if current == target {
					return target
				} else if current > target {
					if current-target < minR {
						minR = current - target
						result = current
					}
					k--
				} else {
					if target-current < minR {
						minR = target - current
						result = current
					}
					j++
				}

			}
		}
	}

	return result
}
