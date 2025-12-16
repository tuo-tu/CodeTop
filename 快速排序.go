package main

import (
	"math/rand"
	"slices"
)

// 核心思想：快速排序（分治 + 递归）
// 选择一个基准值，将数组分为两部分，递归处理
// 时间复杂度：平均 O(n log n)，最坏 O(n²)，空间复杂度：O(log n)
// 在 nums 中随机选择一个基准元素 pivot
// 根据 pivot 重新排列 nums
// 重新排列后，<= pivot 的元素都在 pivot 的左侧，>= pivot 的元素都在 pivot 的右侧
// 返回 pivot 在重新排列后的 nums 中的下标
// 特别地，如果数组的所有元素都等于 pivot，我们会返回数组的中心下标，避免退化
func partition(nums []int) int {
    // 1. 在 nums 中随机选择一个基准元素 pivot
    n := len(nums)
    i := rand.Intn(n)
    pivot := nums[i]
    // 把 pivot 与第一个元素交换，避免 pivot 干扰后续划分，从而简化实现逻辑
    nums[i], nums[0] = nums[0], nums[i]

    i, j := 1, n-1
    for {
        for i <= j && nums[i] < pivot {
            i++
        }
        // 此时 nums[i] >= pivot

        for i <= j && nums[j] > pivot {
            j--
        }
        // 此时 nums[j] <= pivot

        if i >= j {
            break
        }

        // 维持循环不变量
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
    nums[0], nums[j] = nums[j], nums[0]
    return j
}

// 快速排序 nums
func sortArray1(nums []int) []int {
    // 优化：如果 nums 已是升序，直接返回
    if slices.IsSorted(nums) {
        return nums
    }

    i := partition(nums)  // 划分 nums
    sortArray1(nums[:i])   // 排序在 pivot 左侧的元素
    sortArray1(nums[i+1:]) // 排序在 pivot 右侧的元素
    return nums
}