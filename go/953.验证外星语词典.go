/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
logo
探索
题库
讨论
竞赛
企业
商店


验证外星语词典
提交记录
119 / 119 个通过测试用例
状态：通过
执行用时: 0 ms
内存消耗: 2.6 MB
提交时间：几秒前
执行用时分布图表
执行消耗内存分布图表
邀请好友来挑战 验证外星语词典
提交的代码： 几秒前
语言： golang


func isAlienSorted(words []string, order string) bool {
    if len(words) <= 1 {
        return true
    }

    mp := make(map[byte]int)
    for i:=0; i<len(order); i++ {
        mp[order[i]] = i
    }

    for i:=1; i<len(words); i++ {
        l1, l2 := len(words[i-1]), len(words[i])
        var start1, start2 int
        for start1 < l1 && start2 < l2 {
            if mp[words[i-1][start1]] < mp[words[i][start2]] {
                break
            } else if mp[words[i-1][start1]] > mp[words[i][start2]] {
                return false
            } 
            start1++
            start2++
        }

        if start2 == l2 && start1 < l1 {
            return false
        }
    }

    return true
}
// @lc code=end

