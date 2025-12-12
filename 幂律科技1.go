package main

import "fmt"

// 已知数组 A, B, 如果 A 中元素在 B数组存在，打印出这个元素的下标，B 数组是不重复的.
// Input: A = [5, 3, 1, 5, 4]，B = [5, 3]
// Output: [0, 1, 3]

func test1(A, B []int) []int {
	res := []int{}
	hash := map[int]bool{}

	for _, v := range B {
		hash[v] = true
	}

	for i, v := range A {
		if _, ok := hash[v]; ok {
			res = append(res, i)
		}
	}
	return res
}

func Test1Main() {
	A := []int{5, 3, 1, 5, 4}
	B := []int{5, 3}
	fmt.Println(test1(A, B))
}
