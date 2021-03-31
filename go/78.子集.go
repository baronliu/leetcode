/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
// func subsets(nums []int) [][]int {
// 	//初始化结果集
// 	result := make([][]int, 0)
// 	result = append(result, make([]int, 0))
// 	var l int
// 	for i := 0; i < len(nums); i++ {
// 		//当前这个循环内的个数
// 		l = len(result)
// 		for j := 0; j < l; j++ {
// 			info := make([]int, len(result[j]))
// 			copy(info, result[j])
// 			fmt.Println(info)
// 			result = append(result, append(info, nums[i]))
// 		}
// 	}

// 	return result
// }

//解法2
// func subsets(nums []int) [][]int {
//     ret := make([][]int, 0)
//     var stackBack func([]int, int)
//     stackBack = func(cur []int, start int) {
//         ret = append(ret, cur)
//         for i := start; i < len(nums); i++ {
//             tmp := make([]int, len(cur))
//             copy(tmp, cur)
//             tmp = append(tmp, nums[i])
//             stackBack(tmp, i+1)
//         }
//     }
//     stackBack(make([]int, 0), 0)
//     return ret
// }

//解法3 回溯
func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	var stackBack func([]int, int)
	stackBack = func(cur []int, i int) {
		if i == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
			return
		}
		//选择了当前
		cur = append(cur, nums[i])
		stackBack(cur, i+1)
		//不选择当前
		cur = cur[:len(cur)-1]
		stackBack(cur, i+1)
	}

	stackBack(make([]int, 0), 0)

	return result
}

// @lc code=end

