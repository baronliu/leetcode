/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
    var result []int
    if len(nums1) < 1 || len(nums2) < 1 {
        return result
    }

    dict := make(map[int]int)

    for i:=0; i<len(nums1); i++ {
        dict[nums1[i]]++
    }

    for j:=0; j<len(nums2); j++ {
        //如果存在于第一个且剩余个数不为0
        if v := dict[nums2[j]]; v > 0 {
            result = append(result, nums2[j])
            dict[nums2[j]]--
        }
    }

    return result
}
// @lc code=end

