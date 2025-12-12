package main

// 核心思想：滑动窗口 + 哈希表
// 使用双指针 left 和 right 维护一个滑动窗口
// 用哈希表记录每个字符最后出现的位置
// 当遇到重复字符时，将 left 移动到重复字符上次出现位置的下一位
// 在遍历过程中不断更新最大长度
// 时间复杂度：O(n)，空间复杂度：O(min(m,n))，m为字符集大小
func lengthOfLongestSubstring(s string) int {
	// 存储字符串的位置
	buf := make(map[byte]int)
	left := 0
	maxLen := 0

	// 不能用 for range，会导致类型不匹配
	for right := 0; right < len(s); right++ {
		// 存在该字符且idx在窗口里面，则更新left
		if idx, ok := buf[s[right]]; ok && idx >= left {
			// idx此时还未更新
			left = idx + 1
		}

		// 更新字符的位置，
		buf[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
