package main

import "fmt"

// 使用map
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
