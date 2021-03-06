/*
 * @lc app=leetcode.cn id=365 lang=golang
 *
 * [365] 水壶问题
 *
 * https://leetcode-cn.com/problems/water-and-jug-problem/description/
 *
 * algorithms
 * Medium (27.32%)
 * Likes:    157
 * Dislikes: 0
 * Total Accepted:    16.2K
 * Total Submissions: 48.2K
 * Testcase Example:  '3\n5\n4'
 *
 * 有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
 *
 * 如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
 *
 * 你允许：
 *
 *
 * 装满任意一个水壶
 * 清空任意一个水壶
 * 从一个水壶向另外一个水壶倒水，直到装满或者倒空
 *
 *
 * 示例 1: (From the famous "Die Hard" example)
 *
 * 输入: x = 3, y = 5, z = 4
 * 输出: True
 *
 *
 * 示例 2:
 *
 * 输入: x = 2, y = 6, z = 5
 * 输出: False
 *
 *
 */

// @lc code=start
func canMeasureWater(x int, y int, z int) bool {
	if x == 0 || y == 0 {
		return x+y == z || z == 0
	} else if x+y < z {
		return false
	} else if x > y {
		return z%gcd(x, y) == 0
	} else {
		return z%gcd(y, x) == 0
	}
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// @lc code=end

