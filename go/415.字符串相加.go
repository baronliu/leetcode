/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	if l1 == 0 {
		return num2
	}
	if l2 == 0 {
		return num1
	}

	i, j, add := l1-1, l2-1, 0

	ans := ""
	for i >= 0 || j >= 0 {
		sum := add
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		add = sum / 10
		ans = strconv.Itoa(sum%10) + ans
	}
	if add > 0 {
		ans = strconv.Itoa(add) + ans
	}
	return ans
}

// @lc code=end

