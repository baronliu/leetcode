/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 *
 * https://leetcode-cn.com/problems/maximum-swap/description/
 *
 * algorithms
 * Medium (36.10%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 13K
 * Testcase Example:  '2736'
 *
 * 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
 *
 * 示例 1 :
 *
 *
 * 输入: 2736
 * 输出: 7236
 * 解释: 交换数字2和数字7。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: 9973
 * 输出: 9973
 * 解释: 不需要交换。
 *
 *
 * 注意:
 *
 *
 * 给定数字的范围是 [0, 10^8]
 *
 *
 */

// @lc code=start
func maximumSwap(num int) int {
	if num <= 11 {
		return num
	}
	result := 0
	arr := int2Arr(num)
	cache := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	for i := len(arr) - 1; i >= 0; i-- {
		cache[arr[i]] = i
	}

	for i := len(arr) - 1; i >= 0; i-- {
		//扫描大于当前值的元素
		for j := 9; j > arr[i]; j-- {
			idx := cache[j]
			if idx < i && idx >= 0 {
				arr[i], arr[idx] = arr[idx], arr[i]
				goto breakHere
			}
		}
	}

breakHere:
	result = arr2Int(arr)
	return result
}

func arr2Int(arr []int) int {
	var res = 0
	for i := len(arr) - 1; i >= 0; i-- {
		res = res*10 + arr[i]
	}
	return res
}

func int2Arr(n int) []int {
	var result []int
	for n > 0 {
		result = append(result, n%10)
		n = n / 10
	}
	return result
}

// @lc code=end

