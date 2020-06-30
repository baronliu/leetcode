/*
 * @lc app=leetcode.cn id=1343 lang=golang
 *
 * [1343] 大小为 K 且平均值大于等于阈值的子数组数目
 */

// @lc code=start
func numOfSubarrays(arr []int, k int, threshold int) int {
    //求最大值
    sum := k * threshold
    current, res := 0, 0

    //第一个子数组
    for i:=0; i<k; i++ {
        current += arr[i]
    }

    if current >= sum {
        res++
    }

    for i:=k; i<len(arr); i++ {
        current = current + arr[i] - arr[i-k]
        if current >= sum {
            res++
        }
    }
    return res
}
// @lc code=end

