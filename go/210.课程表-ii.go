/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 *
 * https://leetcode-cn.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (47.25%)
 * Likes:    143
 * Dislikes: 0
 * Total Accepted:    23.7K
 * Total Submissions: 49.5K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 n 门课需要选，记为 0 到 n-1。
 *
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
 *
 * 给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
 *
 * 可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
 *
 * 示例 1:
 *
 * 输入: 2, [[1,0]]
 * 输出: [0,1]
 * 解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
 *
 * 示例 2:
 *
 * 输入: 4, [[1,0],[2,0],[3,1],[3,2]]
 * 输出: [0,1,2,3] or [0,2,1,3]
 * 解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
 * 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
 *
 *
 * 说明:
 *
 *
 * 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 *
 *
 * 提示:
 *
 *
 * 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
 * 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
 *
 * 拓扑排序也可以通过 BFS 完成。
 *
 *
 *
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	var result []int
	inDegrees := make(map[int]int)
	outDegrees := make(map[int][]int)

	// 记录下所有节点的入度 和 出度
	for _, v := range prerequisites {
		inDegrees[v[0]]++
		outDegrees[v[1]] = append(outDegrees[v[1]], v[0])
	}

	// 记录下当前入度为0的课程，即可上的课程
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 没有不依赖其它课程的课程，返回空
	if len(queue) == 0 {
		return result
	}

	for len(queue) > 0 {
		// 取出队头，加入结果集
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// 处理队头的出边, 把入度为0的入队
		for i := 0; i < len(outDegrees[node]); i++ {
			out := outDegrees[node][i]
			inDegrees[out]--
			if inDegrees[out] == 0 {
				queue = append(queue, out)
			}
		}
	}

	// 检查是否所有的课程都被安排，如果没有，说明有环，无法安排，返回空
	if len(result) == numCourses {
		return result
	}

	return []int{}
}

// @lc code=end

