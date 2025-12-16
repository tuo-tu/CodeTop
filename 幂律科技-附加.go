package main

func maxProduct(words []string) int {
	res := 0
	n := len(words)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !hasCommonChar(words[i], words[j]) {
				product := len(words[i]) * len(words[j])
				res = max(res, product)
			}
		}
	}
	return res
}

// 判断两个字符串是否有公共字符
func hasCommonChar(a, b string) bool {
	seen := [26]bool{}
	for _, ch := range a {
		seen[ch-'a'] = true
	}
	for _, ch := range b {
		if seen[ch-'a'] {
			return true
		}
	}
	return false
}
