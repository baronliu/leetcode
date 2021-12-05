/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	var b1, b2 [26]int
	sum := 0
	for i := 0; i < len(s); i++ {
		b1[s[i]-'a']++
		b2[goal[i]-'a']++
		if s[i] != goal[i] {
			sum++
		}
	}

	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if b1[index] != b2[index] {
			return false
		}
		if b1[index] > 1 && sum == 0 {
			return true
		}
	}
	return sum == 2
}

// @lc code=end

