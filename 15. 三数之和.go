package main

import "sort"

// 排序 + 双指针
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
				// 当 nums[l] == nums[l+1] 时，说明 nums[l] 和下一个元素 nums[l+1] 是相同的。
				for l < r && nums[l] == nums[l+1] {
					l++
				}
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
