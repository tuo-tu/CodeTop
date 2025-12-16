package main

// 空间优化
func maxSubArray(nums []int) int {
	n := len(nums)
	// 以当前下标 i 为结尾的子数组的最大和
	cur := nums[0]
	maxSum := nums[0]
	for i := 1; i < n; i++ {
		cur = max(nums[i], cur+nums[i])
		maxSum = max(cur, maxSum)
	}
	return maxSum
}
