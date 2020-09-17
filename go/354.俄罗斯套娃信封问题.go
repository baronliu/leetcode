/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(a, b int) bool {
		return envelopes[a][0] < envelopes[b][0] || (envelopes[a][0] == envelopes[b][0] && envelopes[a][1] > envelopes[b][1])
	})

	//求高的lis
	dp := make([]int, len(envelopes))
	var ret int
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				//求当前的dp最大值
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(ret, dp[i])
	}

	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

