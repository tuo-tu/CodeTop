package main

// 核心思想：快速选择算法（Quick Select）
// 第 k 大的元素在排序后的数组中的位置是 len(nums) - k
// 使用快速排序的 partition 思想，但不需要完全排序整个数组
// 根据 partition 返回的位置与目标位置比较：
//   - 如果等于目标位置，直接返回
//   - 如果小于目标位置，在右半部分继续查找
//   - 如果大于目标位置，在左半部分继续查找
// 时间复杂度：平均 O(n)，最坏 O(n²)，空间复杂度：O(1)
func findKthLargest(nums []int, k int) int {
	// 第 k 大的元素的下标
	target := len(nums) - k
	left := 0
	right := len(nums) - 1
	for {
		index := partion(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target { // 说明目标在右边
			left = index + 1
		} else { // 目标在左边
			right = index - 1
		}
	}
}

// 返回分区的分界点下标
func partion(nums []int, left, right int) int {
	// 不要写成div := right，会增加复杂度
	div_value := nums[right]
	i := left
	for j := left; j < right; j++ {
		// 小于才交换，大于等下一次交换
		if nums[j] < div_value { // 注意是比较值，不是比较索引
			nums[i], nums[j] = nums[j], nums[i]
			// 这时候i直接往前推进
			i++
		}
	}
	// 将right放到正确的位置
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
