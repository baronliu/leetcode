/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (57.50%)
 * Likes:    388
 * Dislikes: 0
 * Total Accepted:    87.7K
 * Total Submissions: 142.9K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 示例 1:
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * 说明:
 *
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 *
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
        return nums[0]
    }
	arr := nums[:k]
	for i := (k - 2)/2; i>=0; i-- {
        heapify(arr, i, k)
    }
	for i:=k; i<len(nums); i++ {
		if arr[0] < nums[i] {
			arr[0] = nums[i]
			heapify(arr, 0, k)
		}
	}
	return arr[0]
}

//以i的索引为顶
func heapify(a []int, i int, l int) {
	if i > l {
		return
	}

	min := i
	left := 2*i + 1
	right := 2*i +2

	if left < l && a[left] < a[min]{
		min = left
	}

	if right < l && a[right] < a[min]{
		min = right
	}

	if min != i {
		swap(a, min, i)
		heapify(a, min, l)
	}
}

func swap(a []int, i, j int)  {
	a[i], a[j] = a[j], a[i]
}
// @lc code=end

