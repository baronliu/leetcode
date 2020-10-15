/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */
func multiply(num1 string, num2 string) string {
	//排除0的情况
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	//获取两个字符串长度 及 初始化结果slice
	l1, l2 := len(num1), len(num2)
	result := make([]int, l1+l2)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			//获取当前乘积
			temp := int(num1[l1-1-i]-'0') * int(num2[l2-1-j]-'0')
			//添加个位
			result[i+j] += temp % 10
			//添加十位
			result[i+j+1] += temp / 10
		}
	}
	info := ""
	for i := 0; i < len(result)-1; i++ {
		info = strconv.Itoa(result[i]%10) + info
		result[i+1] += result[i] / 10
	}

	info = strconv.Itoa(result[len(result)-1]) + info
	startIndex := 0
	for i := 0; i < len(info); i++ {
		if info[i] != '0' {
			startIndex = i
			break
		}
	}

	return info[startIndex:]
}

