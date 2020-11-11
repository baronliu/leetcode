/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func maxSubarraySumCircular(A []int) int {
	if len(A) < 1 {
		return 0
	}
	if len(A) == 1 {
		return A[0]
	}
	maxRet, minRet, maxPrev, minPrev := A[0], A[0], A[0], A[0]

	for i := 1; i < len(A); i++ {
		//kanade，求得当前的最大/小值
		maxPrev = A[i] + max(maxPrev, 0)
		minPrev = A[i] + min(minPrev, 0)

		maxRet = max(maxRet, maxPrev)
		minRet = min(minRet, minPrev)
	}

	total := 0
	for i := 0; i < len(A); i++ {
		total += A[i]
	}

	if minRet == total {
		return maxRet
	}

	return max(maxRet, total-minRet)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

