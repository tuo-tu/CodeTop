package main

// 核心思想：使用栈来匹配括号
// 遇到左括号时入栈，遇到右括号时检查栈顶是否匹配
// 如果匹配则弹出栈顶元素，如果不匹配或栈为空则返回 false
// 最后检查栈是否为空，为空说明所有括号都正确匹配
// 时间复杂度：O(n)，空间复杂度：O(n)
func isValid(s string) bool {
	// 定义一个栈来存储左括号。左括号遇到右括号时会弹出，因此栈里面不会存储右括号
	stack := []rune{}
	// 每个右括号对应一个左括号
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	// 遍历字符串
	for _, ch := range s {
		// 先判断是左括号还是右括号
		if ch == ')' || ch == '}' || ch == ']' { // 遇到右括号
			// 已经有了右括号,栈为空和栈顶不匹配都表明栈里面没有和 val 匹配的左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false // 直接结束
			}
			// 如果有匹配到的左括号，弹出
			stack = stack[:len(stack)-1]
		} else { // 遇到左括号
			// 左括号直接入栈
			// 不能写成 stack[len(stack) - 1] = val,会报错
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
