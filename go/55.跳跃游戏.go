/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		if cover >= len(nums)-1 {
			return true
		}
		cover = max(cover, i+nums[i])
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

