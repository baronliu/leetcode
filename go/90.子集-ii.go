/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	//先排序
	sort.Ints(nums)
	ret := make([][]int, 0)
	var stackBack func([]int, int)
	stackBack = func(cur []int, index int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ret = append(ret, tmp)
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			stackBack(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	stackBack([]int{}, 0)
	return ret
}

// @lc code=end

