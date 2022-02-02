	/*
	 * @lc app=leetcode.cn id=2000 lang=golang
	 *
	 * [2000] 反转单词前缀
	 */

	// @lc code=start
	func reversePrefix(word string, ch byte) string {
		ans := ""
		start := 0
		for start < len(word) {
			ans = word[start:start+1] + ans
			if ch == word[start] {
				return ans + word[start+1:]
			}
			start++
		}

		return word
	}

	// @lc code=end

