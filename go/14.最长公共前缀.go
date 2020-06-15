/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (32.17%)
 * Total Accepted:    71.4K
 * Total Submissions: 217.3K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func longestCommonPrefix(strs []string) string {
    //空数组
    if len(strs) < 1 {
        return ""
    }
    //单元素数组结果为其本身
    if len(strs) == 1 {
        return strs[0]
    }
    //两两比较到最后一组
    prev := strs[0]
    for i:=1; i<len(strs); i++ {
        //如果比较而得的元素为空，则停止比较
        if len(prev) == 0 {
            return ""
        }
        prev = compare(prev, strs[i])
    }
    return prev
}

//求两个字符串的公共部分
func compare(str1, str2 string) string {
	l := len(str1)
	if l > len(str2) {
		l = len(str2)
	}
	i := 0
	for i < l {
		if str1[i] != str2[i] {
			break
		}
		i++
	}
	return str1[0:i]
}

