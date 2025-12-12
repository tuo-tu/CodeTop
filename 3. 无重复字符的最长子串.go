package main

// map + 滑动窗口
func lengthOfLongestSubstring(s string) int {
	// 存储字符串的位置
	buf := make(map[byte]int)
	left := 0
	maxLen := 0

	// 不能用 for range，会导致类型不匹配
	for right := 0; right < len(s); right++ {
		// idx在窗口里面
		if idx, ok := buf[s[right]]; ok && idx >= left {
			// idx此时还未更新
			left = idx + 1
		}

		// 更新字符的位置
		buf[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
