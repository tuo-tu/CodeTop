package main

import "fmt"

// 核心思想：使用哈希表实现一次遍历
// 遍历数组时，用哈希表记录已访问过的数字及其索引
// 对于当前数字 nums[i]，检查 target - nums[i] 是否在哈希表中
// 如果存在，说明找到了两个数，返回它们的索引
// 时间复杂度：O(n)，空间复杂度：O(n)
func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := hashmap[target-nums[i]]; ok {
			return []int{i, j}
		}
		// 不能写成 hashmap[target-nums[i]] = i
		hashmap[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 3, 5, 6, 7, 7, 6}, 9))
}
