/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
//x & (x-1) 会消除掉最后一个二进制的1
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = countOnes(i)
	}
	return ans
}

func countOnes(x int) int {
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

// @lc code=end

