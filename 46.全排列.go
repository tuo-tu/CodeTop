package main

// 核心思想：回溯法 + 交换
// 通过交换元素来生成所有可能的排列
// 对于位置 start，枚举 [start..len-1] 中的每个元素，将其与 start 位置交换
// 然后递归处理 start+1 位置，递归返回后需要回溯（恢复交换）
// 当 start 到达数组末尾时，说明得到一个完整的排列
// 时间复杂度：O(n × n!)，空间复杂度：O(n)（递归栈深度）
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
