/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	length := len(s)
	end := length - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return 0
	}

	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return end - start
}

// @lc code=end

