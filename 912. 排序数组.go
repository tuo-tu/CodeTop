package main

// 归并排序
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	// 开始递归
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	return merge1(left, right)
}

func merge1(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
		// i++ ,j++写到这里是错的
	}
	// 处理剩下的元素,左右一定有一边是空的，注意要放在循环之外
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

