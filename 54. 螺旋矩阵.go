package main

import "math"

// 方向数组 + 原地标记
// 始终沿着当前方向前进；如果下一个格子越界或已访问，就右转；走满 m*n 步就结束。
func spiralOrder(matrix [][]int) []int {
	// 特殊情况
	if len(matrix) == 0 {
		return []int{}
	}
	// 分别代表行数、列数
	m, n := len(matrix), len(matrix[0])
	// 结果集
	res := make([]int, 0, m*n)
	// 定义 4 个方向：右、下、左、上
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 从左上角 (0,0) 开始，方向（0=右，1=下，2=左，3=上）
	x, y, d := 0, 0, 0
	v := math.MinInt // 用于填充每个走过的地方
	// 开始处理所有节点
	for i := 0; i < m*n; i++ {
		res = append(res, matrix[x][y]) // 把当前节点值加到结果中
		matrix[x][y] = v                // 标记已经走过
		xn, yn := x+dirs[d][0], y+dirs[d][1]
		// 判断是否越界或遍历过,判断xn、yn而不是x
		if xn < 0 || xn >= m || yn < 0 || yn >= n || matrix[xn][yn] == v {
			// 转向
			d = (d + 1) % 4
			// 重新赋值
			xn, yn = x+dirs[d][0], y+dirs[d][1]
		}
		// 下一个节点
		x, y = xn, yn
	}
	return res
}
