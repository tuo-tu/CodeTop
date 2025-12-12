package main

import (
	"fmt"
	"math"
)

// 从上到下找到最短路径(n个数字之和最小,n为矩阵的行数)，
// 可以从第一行中的任何元素开始，只能往下层走同时只能走向相邻的节点，
// 例如图中第一排 2 只能走向 第二排的 7、3;第二排的 7 可以走向第三排的 6、2、9
//
// |5|8|1|2|
// |4|1|7|3|
// |3|6|2|9|
//
// Input: [
//	[5,8,1,2],
//	[4,1,7,3],
//	[3,6,2,9]
// ]
// Output: 4

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// m,n 分别代表行数和列数
	row, col := len(matrix), len(matrix[0])
	// dp[i][j]代表自上而下的节点的最短路径
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	// 第一行沿用 matrix 的数字
	for i := 0; i < col; i++ {
		dp[0][i] = matrix[0][i]
	}

	// 开始填充剩下的位置
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			minPre := math.MaxInt
			// 左上方的
			if j > 0 {
				minPre = min(minPre, dp[i-1][j-1])
			}
			// 正上方的
			minPre = min(minPre, dp[i-1][j])
			// 右上方的
			if j+1 < col {
				minPre = min(minPre, dp[i-1][j+1])
			}
			// 不要忘记赋值给dp[i][j]
			dp[i][j] = matrix[i][j] + minPre
		}
	}
	// 直接获取最后一行的结果即可
	minPath := math.MaxInt
	for i := 0; i < col; i++ {
		minPath = min(minPath, dp[row-1][i])
	}
	return minPath
}

func test3(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = matrix[0][i]
	}

	// minPre := math.MaxInt
	for i := 1; i < m; i++ {

		for j := 0; j < n; j++ {
			minPre := math.MaxInt
			if j > 0 {
				minPre = min(minPre, dp[i-1][j-1])
			}
			minPre = min(minPre, dp[i-1][j])
			if j+1 < n {
				minPre = min(minPre, dp[i-1][j+1])
			}

			dp[i][j] = dp[i][j] + minPre
		}
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dp[m-1][i])
	}
	return res
}

func Test3Main() {
	matrix := [][]int{
		{5, 8, 1, 2},
		{4, 1, 7, 3},
		{3, 6, 2, 9},
	}
	//fmt.Println(minFallingPathSum(matrix))
	fmt.Println(test3(matrix))
}
