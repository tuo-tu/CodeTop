package main

import (
	"sort"
)

// 动态规划：单切片
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

// 二分查找
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
