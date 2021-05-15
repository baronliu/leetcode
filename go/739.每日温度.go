/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start

func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	if len(T) == 1 {
		return ans
	}

	stack := make([]int, 0)

	for k, v := range T {
		for len(stack) > 0 && v > T[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = k - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, k)
	}

	return ans
}

// @lc code=end

