/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	nums := make([]int, 0)
	backtrack(0, target, nums, candidates, &ans)
	return ans
}

func backtrack(index, target int, nums, candidates []int, ans *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		ret := make([]int, len(nums))
		copy(ret, nums)
		*ans = append(*ans, ret)
		return
	}

	for i := index; i < len(candidates); i++ {
		//向前
		target = target - candidates[i]
		nums = append(nums, candidates[i])
		backtrack(i, target, nums, candidates, ans)
		//回退
		target = target + candidates[i]
		nums = nums[:len(nums)-1]
	}
}

// @lc code=end

