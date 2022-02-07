/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}
	left, size := 0, 0
	base, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len1; i++ {
		base[s1[i]]++
	}

	for right := 0; right < len2; right++ {
		c := s2[right]
		if v, ok := base[c]; ok {
			window[c]++
			if v == window[c] {
				size++
			}
		}

		if right-left+1 < len1 {
			continue
		}

		if size == len(base) {
			return true
		}

		c = s2[left]
		if v, ok := base[c]; ok {
			if v == window[c] {
				size--
			}
			window[c]--
		}
		left++
	}
	return false
}

// @lc code=end

