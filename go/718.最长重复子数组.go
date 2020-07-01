/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(A []int, B []int) int {
    max := 0
    dp := make([][]int, len(A))
    for i:=0; i<len(dp); i++ {
        dp[i] = make([]int, len(B))
    }

    for i:=0; i<len(A); i++ {
        dp[i] = make([]int, len(B))
        for j:=0; j<len(B); j++ {
            if A[i] == B[j] {
                if i==0 || j==0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = dp[i-1][j-1] + 1
                }
            } else {
                dp[i][j] = 0
            }

            if dp[i][j] > max {
                max = dp[i][j]
            }
        }
    }

    return max
}
// @lc code=end

