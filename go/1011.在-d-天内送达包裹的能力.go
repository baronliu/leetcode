/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, weight := range weights {
		right += weight
		if weight > left {
			left = weight
		}
	}

	for left < right {
		mid := (right + left) / 2
		if check(weights, D, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(weights []int, D, C int) bool {
	cur := 0
	for i := 0; i < len(weights); i++ {
		if cur+weights[i] > C {
			if D <= 1 {
				return false
			}
			D--
			cur = 0

		}
		cur += weights[i]
	}
	return true
}

// @lc code=end

