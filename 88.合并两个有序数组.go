package main

// 逆向双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	p := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// 注意 nums1 比 nums2 长
	// 因此最后只有可能nums2的剩余元素需要被处理，
	// 如果nums1剩余，不需要被处理
	// 如果 nums2 还有剩余，直接拷贝过去
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}


func mergeSortedArrays(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	result := make([]int, 0, m+n)

	// 双指针遍历
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	// 把剩余的元素追加上
	if i < m {
		result = append(result, nums1[i:]...)
	}

	if j < n {
		result = append(result, nums2[j:]...)
	}

	return result
}
