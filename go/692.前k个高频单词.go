/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
	newWords := make([]string, 0)
	dict := make(map[string]int)
	for _, word := range words {
		dict[word]++
		if dict[word] == 1 {
			newWords = append(newWords, word)
		}
	}

	sort.Slice(newWords, func (i, j int) bool {
		return (dict[newWords[i]] > dict[newWords[j]]) || (dict[newWords[i]] == dict[newWords[j]] && newWords[i] < newWords[j]) 
	})

	return newWords[:k]
}

// @lc code=end

