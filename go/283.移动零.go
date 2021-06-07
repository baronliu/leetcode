/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	length := len(nums)
	var index int
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < length; i++ {
		nums[i] = 0
	}
}

// @lc code=end

