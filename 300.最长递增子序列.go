package main

import (
	"sort"
)

// 核心思想：动态规划
// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
// 对于每个位置 i，遍历前面所有位置 j，如果 nums[j] < nums[i]，
// 则可以将 nums[i] 接在以 nums[j] 结尾的子序列后面，更新 dp[i]
// 时间复杂度：O(n²)，空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	// 特殊情况
	if n == 0 {
		return 0
	}
	// dp[i] 存储以 nums[i] 结尾的字符串的严格递增子序列长度
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		// 长度至少是1
		dp[i] = 1
	}
	// 记录结果，长度至少是1
	res := 1
	// 开始处理字符串
	for i := 1; i < n; i++ {
		// 处理从0到i的位置的字符串
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				// 取较长的一个
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// 核心思想：贪心 + 二分查找（使用 sort.SearchInts）
// tails[i] 表示长度为 i+1 的递增子序列的最小末尾元素
// 对于每个 num，在 tails 中二分查找第一个 >= num 的位置
// 如果找到，替换该位置的值（保证结尾更小，有利于后续扩展）
// 如果没找到（num 比所有元素都大），追加到 tails 末尾
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func lengthOfLIS1(nums []int) int {
	tails := []int{}
	for _, num := range nums {
		// 找到第一个 >= num 的位置,当tails空的时候，i返回0
		i := sort.SearchInts(tails, num)
		// 表示 num 比 tails 里所有数都大
		if i == len(tails) {
			tails = append(tails, num)
		} else {
			tails[i] = num // 替换，保证结尾更小
		}
	}
	return len(tails)
}

// 核心思想：贪心 + 二分查找（手动实现二分）
// tails[i] 表示长度为 i+1 的递增子序列的最小末尾元素
// 对于每个 num，手动二分查找 tails 中第一个 >= num 的位置
// 如果找到，替换该位置的值；如果没找到，追加到 tails 末尾
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func lengthOfLIS2(nums []int) int {
	// tails[i] 等于长度为 i+1 的最小 nums[i] 的值
	tails := []int{}
	// 开始处理 nums，目标是找到 tails 中第一个 >= num 的位置
	for _, num := range nums {
		l, r := 0, len(tails)
		// 这个循环的目的是获取最终的l
		for l < r {
			mid := (l + r) / 2
			if tails[mid] < num { // mid 及其左边的元素都不可能 >= num
				l = mid + 1
			} else { // 当 tails[mid] >= num，说明可能答案就在 mid 或 mid 左边
				r = mid
			}
		}

		// 最终循环结束时：l == r → 第一个 >= num 的位置
		// 如果 l 等于 tails 长度，说明 num 比所有结尾都大，扩展序列
		// 当 tails 为空，i 返回 0
		if l == len(tails) {
			tails = append(tails, num)
		} else {
			// 否则，替换最小的 >= num 的位置
			tails[l] = num
		}

	}
	return len(tails)
}
