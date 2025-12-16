package main

func firstUniqChar(s string) int {
	// 特殊情况处理
	if len(s) == 0 {
		return -1
	}

	// 存储 s 中的每个字符出现的次数
	count := map[rune]int{}
	for _, ch := range s {
		// 一开始 count 是空的，但是 map 会自动分配一个零值
		count[ch]++
	}

	// 找第一个只出现一次的字符
	for i, ch := range s {
		if count[ch] == 1 {
			return i
		}
	}
	return -1
}
