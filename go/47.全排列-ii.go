/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)

	if len(nums) < 1 {
		return ans
	}

	if len(nums) == 1 {
		ans = append(ans, []int{nums[0]})
	}

	used := make(map[int]interface{})
	for i := 0; i < len(nums); i++ {
		//保存掉已计算过的尾部
		if _, ok := used[nums[i]]; !ok {
			used[nums[i]] = nil
			tmp := make([]int, 0)
			for _, v := range nums[:i] {
				tmp = append(tmp, v)
			}
			for _, v := range nums[i+1:] {
				tmp = append(tmp, v)
			}
			cur := permuteUnique(tmp)
			for _, v := range cur {
				v = append(v, nums[i])
				ans = append(ans, v)
			}
		}
	}
	return ans
}

// @lc code=end

