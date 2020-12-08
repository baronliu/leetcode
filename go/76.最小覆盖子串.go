/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	base := make(map[byte]int)
	for i := range t {
		base[t[i]]++
	}

	var left, right, size int
	start, length := 0, math.MaxInt32
	window := make(map[byte]int)
	for right < len(s) {
		tempRight := s[right]
		right++
		if _, ok := base[tempRight]; ok {
			window[tempRight]++
			//满足了其中一项
			if window[tempRight] == base[tempRight] {
				size++
			}
		}

		for size == len(base) {
			//满足条件的情况
			if right-left < length {
				length = right - left
				start = left
			}
			tempLeft := s[left]
			left++
			if _, ok := base[tempLeft]; ok {
				window[tempLeft]--
				//满足了其中一项
				if window[tempLeft] < base[tempLeft] {
					size--
				}
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length]
}

// @lc code=end

