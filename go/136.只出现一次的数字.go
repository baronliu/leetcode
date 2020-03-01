/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

package leetcode

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	result := nums[len(nums)-1]
	for i < len(nums)-1 {
		if nums[i] != nums[i+1] {
			result = nums[i]
			break
		} else {
			i = i + 2
		}
	}

	return result
}
