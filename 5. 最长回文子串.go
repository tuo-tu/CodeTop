package main

// 中心扩展法
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	// 定义两组相对应的变量
	start, end := 0, 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		// 奇数长度的回文处理
		left, right = expand(s, i, i)
		if right-left > end-start {
			start, end = left, right
		}
		// 偶数长度的回文处理
		left, right = expand(s, i, i+1)
		if right-left > end-start {
			start, end = left, right
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 返回正确的坐标
	return left + 1, right - 1
}
