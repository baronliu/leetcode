/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] 字典序的第K小数字
 *
 * https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (28.51%)
 * Likes:    71
 * Dislikes: 0
 * Total Accepted:    2.8K
 * Total Submissions: 8.8K
 * Testcase Example:  '13\n2'
 *
 * 给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。
 *
 * 注意：1 ≤ k ≤ n ≤ 10^9。
 *
 * 示例 :
 *
 *
 * 输入:
 * n: 13   k: 2
 *
 * 输出:
 * 10
 *
 * 解释:
 * 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
 *
 *
 */

// @lc code=start
func findKthNumber(n int, k int) int {
	p, prefix := 1, 1
	for p < k {
		count := getCount(prefix, n)
		if p+count > k {
			//缩小范围
			prefix *= 10
			p += 1
		} else {
			prefix += 1
			p = p + count
		}
	}
	return prefix
}

//求一个前缀下的总数
func getCount(prefix int, n int) int {
	count := 0
	for a, b := prefix, prefix+1; a <= n; a, b = a*10, b*10 {
		//如果在这个区间内
		count += min(b, n+1) - a
	}
	return count
}

func min(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

// @lc code=end

