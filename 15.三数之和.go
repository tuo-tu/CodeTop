package main

import "sort"

// 核心思想：排序 + 双指针
// 首先对数组排序，然后固定第一个数，使用双指针在剩余数组中寻找另外两个数
// 固定 nums[i]，在 [i+1, n-1] 区间内使用双指针 left 和 right
// 根据三数之和与 0 的大小关系移动指针：和小于0则left右移，和大于0则right左移
// 注意去重：跳过重复的元素，避免结果集中出现重复的三元组
// 时间复杂度：O(n²)，空间复杂度：O(1)（不考虑结果数组）
func threeSum(nums []int) [][]int {
	// 首先进行排序
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}
	// i < n-2是保证 i 的右边还有两个数可以凑成3个数
	for i := 0; i < n-2; i++ {
		// 注意去重。i > 0 是为了防止越界，因为后面有 i-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				// 注意去重，缩小范围
				// 当 nums[l] == nums[l+1] 时，说明 nums[l] 和下一个元素 nums[l+1] 是相同的,则left右移
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 当 nums[r] == nums[r-1] 时，说明 nums[r] 和前一个元素 nums[r-1] 是相同的,则right左移
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
