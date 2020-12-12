/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	start := len(nums) - 2
	for start >= 0 {
		if nums[start] < nums[start+1] {
			//记住在这个节点
			break
		}
		start--
	}
	if start < 0 {
		sort.Ints(nums)
		return
	}

	//2.找到第一个比start大的值，并交换
	for i := len(nums) - 1; i > start; i-- {
		if nums[i] > nums[start] {
			nums[i], nums[start] = nums[start], nums[i]
			break
		}
	}

	//3.重置掉后面的内容
	sort.Ints(nums[start+1:])
}

// @lc code=end

