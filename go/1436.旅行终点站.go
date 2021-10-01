/*
 * @lc app=leetcode.cn id=1436 lang=golang
 *
 * [1436] 旅行终点站
 */

// @lc code=start
func destCity(paths [][]string) string {
	dict := make(map[string]interface{})
	for _, v := range paths {
		dict[v[0]] = nil
	}

	for _, v := range paths {
		if _, ok := dict[v[1]]; !ok {
			return v[1]
		}
	}
	return ""
}

// @lc code=end

