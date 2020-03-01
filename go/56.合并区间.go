import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (36.93%)
 * Likes:    293
 * Dislikes: 0
 * Total Accepted:    55.2K
 * Total Submissions: 136.6K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	var result [][]int
	if len(intervals) == 0 {
		return result
	}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		tmp := intervals[i]
		for i+1 < len(intervals) && intervals[i+1][0] <= tmp[1] {
			if intervals[i+1][1] > tmp[1] {
				tmp[1] = intervals[i+1][1]
			}
			i++
		}
		result = append(result, tmp)
	}

	return result
}

// @lc code=end

