/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return findIndex(nums, mid)
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

func findIndex(nums []int, current int) []int {
	left, right := current, current
	for left >= 0 && nums[left] == nums[current] {
		left--
	}

	for right <= len(nums)-1 && nums[right] == nums[current] {
		right++
	}
	return []int{left + 1, right - 1}
}

// @lc code=end

