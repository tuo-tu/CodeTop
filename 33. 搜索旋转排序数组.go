package main

// for循环 + 二分查找.
// 注意：target是目标值，而返回的是下标
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 开启循环，以左右边界为条件
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 一定注意，一定有一边严格递增，左或者右。
		// 这里不能用 <,一定要用 <=
		if nums[left] <= nums[mid] { // 只能说明左半边严格递增
			// 再判断 target 是否在左半边
			if target >= nums[left] && target < nums[mid] { // target 在左半边
				// 缩小左边范围
				right = mid - 1
			} else { // target 在右半边，跳到右半边找
				// 注意不能写成left++
				left = mid + 1
			}
		} else { // 只能说明右边严格递增了
			if target > nums[mid] && target <= nums[right] { // target 在右半边
				left = mid + 1
			} else { // target 在左半边，跳到左半边找
				right = mid - 1
			}
		}
	}
	return -1
}
