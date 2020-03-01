/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] Fibonacci Number
 */
func fib(N int) int {
    if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	n1 := 0
	n2 := 1

	for index := 2; index <= N; index++ {
		temp := n1 + n2
		n1 = n2
		n2 = temp
	}

	return n2
}

