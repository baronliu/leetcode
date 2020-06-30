/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
    if len(digits) < 1 {
        return nil
    }

    dict := map[byte][]byte{
        '2': {'a', 'b', 'c'},
        '3': {'d', 'e', 'f'},
        '4': {'g', 'h', 'i'},
        '5': {'j', 'k', 'l'},
        '6': {'m', 'n', 'o'},
        '7': {'p', 'q', 'r', 's'},
        '8': {'t', 'u', 'v'},
        '9': {'w', 'x', 'y', 'z'},
    }
    result := make([]string, 1)

    for i:=0; i<len(digits); i++ {
        digit := digits[i]
        current := dict[digit]
        l := len(result)
        for _, v := range current {
            for j:=0; j<l; j++ {
                result = append(result, string(result[j]) + string(v))
            }
        }
        result = result[l:]
    }

    return result
}
// @lc code=end

