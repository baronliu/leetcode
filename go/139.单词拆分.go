/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
//1.dfs + 记忆
func wordBreak(s string, wordDict []string) bool {

	wordSet := make(map[string]bool)
	for _, v := range wordDict {
		wordSet[v] = true
	}
	memo := make(map[string]bool)

	var dfs func(string) bool
	dfs = func(cur string) bool {
		if cur == "" {
			return true
		}

		if v, ok := memo[cur]; ok {
			return v
		}

		for i := 0; i < len(cur); i++ {
			if wordSet[cur[:i+1]] && dfs(cur[i+1:]) {
				memo[cur] = true
				return true
			}
		}
		memo[cur] = false
		return false
	}

	return dfs(s)
}

//2.动态规划
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, v := range wordDict {
		wordSet[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end

