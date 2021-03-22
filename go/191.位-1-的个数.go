/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var ans int

	for num != 0 {
		if num&1 > 0 {
			ans++
		}
		num = num >> 1
	}

	return ans
}

// @lc code=end

