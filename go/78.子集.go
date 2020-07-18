/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	//初始化结果集
	result := make([][]int, 0)
	result = append(result, make([]int, 0))
	var l int
	for i := 0; i < len(nums); i++ {
		//当前这个循环内的个数
		l = len(result)
		for j := 0; j < l; j++ {
			info := make([]int, len(result[j]))
			copy(info, result[j])
			fmt.Println(info)
			result = append(result, append(info, nums[i]))
		}
	}

	return result
}
// @lc code=end

