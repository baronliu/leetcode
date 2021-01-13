/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
// func permute(nums []int) [][]int {
// 	var ans [][]int
// 	if len(nums) < 1 {
// 		return ans
// 	}
// 	cur := make([]int, 0)
// 	used := make([]bool, len(nums))
// 	backtrack(nums, cur, used, &ans)
// 	return ans
// }

// func backtrack(nums, cur []int, used []bool, ans *[][]int) {
// 	if len(cur) == len(nums) {
// 		tmp := make([]int, len(nums))
// 		copy(tmp, cur)
// 		*ans = append(*ans, tmp)
// 		return
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if !used[i] {
// 			used[i] = true
// 			cur = append(cur, nums[i])
// 			backtrack(nums, cur, used, ans)
// 			used[i] = false
// 			cur = cur[:len(cur)-1]
// 		}
// 	}
// }

func permute(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 1 {
		return ans
	}

	if len(nums) == 1 {
		ans = append(ans, []int{nums[0]})
	}

	for i := 0; i < len(nums); i++ {
		remain := make([]int, 0)
		for _, v := range nums[:i] {
			remain = append(remain, v)
		}
		for _, v := range nums[i+1:] {
			remain = append(remain, v)
		}
		cur := permute(remain)
		for _, v := range cur {
			v = append(v, nums[i])
			ans = append(ans, v)
		}
	}

	return ans
}

// @lc code=end

