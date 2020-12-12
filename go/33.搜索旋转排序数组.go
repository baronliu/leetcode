/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if target >= nums[left] && target <= nums[mid] {
			right = mid
		} else if target <= nums[right] && target > nums[mid] {
			left = mid + 1
		} else if nums[mid] >= nums[left] && (target <= nums[right] || target > nums[mid]) {
			left = mid + 1
		} else if nums[mid] <= nums[right] && (target <= nums[mid] || target >= nums[left]) {
			right = mid
		} else {
			return -1
		}
	}

	if nums[left] == target {
		return left
	}

	if nums[right] == target {
		return right
	}

	return -1
}

// @lc code=end

