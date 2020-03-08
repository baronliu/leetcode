import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (33.46%)
 * Likes:    416
 * Dislikes: 0
 * Total Accepted:    49.1K
 * Total Submissions: 129.4K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

// @lc code=start
var result int

func coinChange(coins []int, amount int) int {
	result = math.MaxInt64
	sort.Ints(coins)
	dfs(coins, amount, len(coins)-1, 0)
	if result == math.MaxInt64 {
		return -1
	}

	return result
}

func dfs(coins []int, amount int, index int, count int) {
	if index < 0 {
		return
	}

	for i := amount / coins[index]; i >= 0; i-- {
		nCount := count + i

		//剪枝
		if nCount >= result {
			break
		}

		if amount < coins[index]*i {
			break
		}

		if amount == coins[index]*i {
			result = nCount
			break
		}

		dfs(coins, amount-coins[index]*i, index-1, nCount)
	}
}

// @lc code=end

