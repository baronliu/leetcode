/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (44.77%)
 * Likes:    887
 * Dislikes: 0
 * Total Accepted:    57.2K
 * Total Submissions: 117.1K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */

// @lc code=start
func trap(height []int) int {
	result := 0
	//小于三列的时候存储量为0
	if len(height) < 3 {
		return result
	}

	//1.暴力法(缺点：在找寻最高点的时候浪费了太多比较)
	// for i := 1; i < len(height)-1; i++ {
	// 	left, right, min, leftMax, rightMax := i-1, i+1, 0, 0, 0
	// 	for left >= 0 {
	// 		if height[left] > leftMax {
	// 			leftMax = height[left]
	// 		}
	// 		left--
	// 	}
	// 	for right <= len(height)-1 {
	// 		if height[right] > rightMax {
	// 			rightMax = height[right]
	// 		}
	// 		right++
	// 	}
	// 	if leftMax >= rightMax {
	// 		min = rightMax
	// 	} else {
	// 		min = leftMax
	// 	}

	// 	if min > height[i] {
	// 		result += min - height[i]
	// 	}
	// }

	//2.暴力法优化（在遍历过程中前一次求最大值的结果不用反复再比较）
	leftPrevMax := 0
	rightMaxArr := make([]int, len(height))
	rightMaxArr[len(height)-1] = height[len(height)-1]
	for j := len(height) - 2; j > 1; j-- {
		if rightMaxArr[j+1] >= height[j] {
			rightMaxArr[j] = rightMaxArr[j+1]
		} else {
			rightMaxArr[j] = height[j]
		}
	}

	for i := 1; i < len(height)-1; i++ {
		left, right, min, leftMax, rightMax := i-1, i+1, 0, 0, 0
		if height[left] > leftPrevMax {
			leftMax = height[left]
		} else {
			leftMax = leftPrevMax
		}
		leftPrevMax = leftMax
		rightMax = rightMaxArr[right]

		if leftMax >= rightMax {
			min = rightMax
		} else {
			min = leftMax
		}

		if min > height[i] {
			result += min - height[i]
		}
	}
	return result
}

// @lc code=end

