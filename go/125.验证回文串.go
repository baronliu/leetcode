/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < len(s) && !is(s[left]) {
			left++
		}

		for right >= 0 && !is(s[right]) {
			right--
		}
		if left < right && !equal(s[left], s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func equal(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a = 'a' + a - 'A'
	}

	if b >= 'A' && b <= 'Z' {
		b = 'a' + b - 'A'
	}

	return a == b
}

func is(a byte) bool {
	return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

// @lc code=end

