package main

// 回溯法，需要使用一个递归函数
func permute(nums []int) [][]int {
	// 存储最终结果
	res := [][]int{}
	// 回溯递归函数
	var dfs func(int)
	dfs = func(start int) {
		// 递归终止条件：当 start 走到数组末尾，代表得到了一组结果
		if len(nums) == start {
			buf := make([]int, len(nums))
			copy(buf, nums)
			res = append(res, buf)
		}

		// 枚举：把 [start..len-1] 这些“尚未固定”的元素，依次和 start 的位置交换
		for i := start; i < len(nums); i++ {
			// 交换，注意一开始的 start和 i是相等的，相当于交换自身
			nums[start], nums[i] = nums[i], nums[start]
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start] // 回溯
		}
	}
	dfs(0)
	return res
}
