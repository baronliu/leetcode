/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for i, num := range nums {
		if j, ok := dict[num]; ok {
			if i-j <= k {
				return true
			}
		}
		dict[num] = i
	}
	return false
}

// @lc code=end

