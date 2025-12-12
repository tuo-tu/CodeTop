package main

// 快速排序 + for循环
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

// 返回分区的分分界点下标
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
