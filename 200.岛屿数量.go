package main

// 深度优先（DFS）+ 递归
// 不断淹没已经访问过的岛屿
func numIslands(grid [][]byte) int {
	// 特殊情况
	if len(grid) == 0 {
		return 0
	}

	count := 0
	// m,n 分别用于表示岛的宽、长
	m, n := len(grid), len(grid[0])
	// 每次统计到一个岛（值为1），就递归淹没整个岛屿
	// 注意不能直接写成 dfs := func(i, j int) {}，原因是 Go 的匿名函数在声明自身时不能直接递归
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 控制边界
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			// 碰到边界直接结束
			return
		}
		// 淹没当前位置，标记为'0'
		grid[i][j] = '0'
		// 当前地点的上下左右也要淹没
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 发现了岛屿
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}
