/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	//利用异或的特性：x^x = 0 x^0 = x
	temp := 0
	for _, num := range nums {
		temp ^= num
	}
	//temp为两个元素的异或结果
	index := 0
	for index < 32 {
		if temp&1 > 0 {
			break
		}
		temp = temp >> 1
		index++
	}

	ans := make([]int, 2)

	for _, num := range nums {
		if (num>>index)&1 > 0 {
			ans[0] ^= num
		} else {
			ans[1] ^= num
		}
	}
	return ans
}

// @lc code=end

