/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	// arr := make([]int, len(nums))
	// for _, v := range nums {
	// 	arr[v-1] += 1
	// }
	// ans := make([]int, 0)
	// for k, v := range arr {
	// 	if v == 0 {
	// 		ans = append(ans, k+1)
	// 	}
	// }
	// return ans

	ans := make([]int, 0)
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] < 0 {
			continue
		} else {
			nums[v-1] = -nums[v-1]
		}
	}

	for k, v := range nums {
		if v > 0 {
			ans = append(ans, k+1)
		}
	}

	return ans
}

// @lc code=end

