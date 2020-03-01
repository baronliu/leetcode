/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
func findMin(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	min := nums[0]

    for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			break
		}
	}
	return min
}

