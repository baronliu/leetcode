/*
 * @lc app=leetcode.cn id=1006 lang=golang
 *
 * [1006] 笨阶乘
 */

// @lc code=start
func clumsy(N int) int {
	var ans, i int
	for i < N {
		if i%4 == 0 {
			var cur int
			if i < N-2 {
				cur = (N - i) * (N - i - 1) / (N - i - 2)
			} else if i == N-2 {
				cur = (N - i) * (N - i - 1)
			} else {
				cur = N - i
			}

			if i == 0 {
				ans = cur
			} else {
				ans -= cur
			}
			i += 3
		} else {
			if i%4 == 3 {
				ans += N - i
				i += 1
			}
		}
	}
	return ans
}

// @lc code=end

