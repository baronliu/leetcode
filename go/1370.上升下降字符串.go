/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	if len(s) <= 1 {
		return s
	}
	ret := make([]byte, 0)
	dict := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for len(ret) < len(s) {
		for i := byte('a'); i <= 'z'; i++ {
			if dict[i] > 0 {
				ret = append(ret, i)
				dict[i]--
			}
		}
		for i := byte('z'); i >= 'a'; i-- {
			if dict[i] > 0 {
				ret = append(ret, i)
				dict[i]--
			}
		}
	}

	return string(ret)
}

// @lc code=end

