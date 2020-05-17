/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (46.80%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    36.9K
 * Total Submissions: 72.7K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
 *
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
 *
 * 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
 *
 *
 *
 * 示例 1:
 *
 * 输入: 2, [[1,0]]
 * 输出: true
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
 *
 * 示例 2:
 *
 * 输入: 2, [[1,0],[0,1]]
 * 输出: false
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 * 1 <= numCourses <= 10^5
 *
 *
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegrees := make(map[int]int)
	outDegrees := make(map[int][]int)

	for _, v := range prerequisites {
		//入度增加
		inDegrees[v[0]]++
		//出度增加
		outDegrees[v[1]] = append(outDegrees[v[1]], v[0])
	}
	//可选课程
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		return false
	}

	count := 0

	for len(queue) > 0 {
		//先学当前的课程
		node := queue[0]
		queue = queue[1:]
		count++
		//再学与之关联的课程
		outDegree := outDegrees[node]
		for _, v := range outDegree {
			inDegrees[v]--
			if inDegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if count == numCourses {
		return true
	}

	return false
}

// @lc code=end

