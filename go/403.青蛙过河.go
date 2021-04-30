/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
func canCross(stones []int) bool {
	if stones[1] > 1 {
		return false
	}

	dict := make(map[int]bool)
	for _, stone := range stones {
		dict[stone] = true
	}

	cache := make(map[string]bool)

	var dfs func(int, int) bool

	dfs = func(cur, prev int) bool {
		if !dict[cur] {
			return false
		}

		//有缓存以缓存为准
		if v, ok := cache[strconv.Itoa(cur)+"-"+strconv.Itoa(prev)]; ok {
			return v
		}

		//当前跳到了最后一块石头上
		if cur == stones[len(stones)-1] {
			return true
		}

		ans := false
		for i := prev - 1; i <= prev+1; i++ {
			if i == 0 {
				continue
			}
			curV := dfs(cur+i, i)
			//写入缓存
			cache[strconv.Itoa(cur+i)+"-"+strconv.Itoa(i)] = curV
			ans = ans || curV
		}

		return ans
	}

	return dfs(1, 1)
}

// @lc code=end

