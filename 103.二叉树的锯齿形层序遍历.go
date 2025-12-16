package main

// DFS + 就地头插/尾插
func zigzagLevelOrder(node *TreeNode) [][]int {
	// 存储结果
	res := [][]int{}
	// 创建dfs递归函数
	var dfs func(*TreeNode, int)
	// level 代表当前是第几层
	dfs = func(node *TreeNode, level int) {
		// 递归结束条件
		if node == nil {
			return
		}

		// 首次到达该层，需要新创建一个切片[]int{}
		if len(res) == level {
			res = append(res, []int{})
		}
		if level%2 == 0 { // 偶数层(0,2,4...) 左->右: 尾插
			res[level] = append(res[level], node.Val)
		} else { // 奇数层(1,3,5...) 右->左: 头插
			res[level] = append([]int{node.Val}, res[level]...)
		}

		// 递归处理左右子树
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(node, 0)
	return res
}

// 层序遍历（BFS）+ 控制方向
func zigzagLevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	// 存储每一层的所有节点
	queue := []*TreeNode{root}
	// 第一层（根）从左到右。
	leftToRight := true

	for len(queue) > 0 {
		size := len(queue)
		// 预分配当前层的结果切片
		level := make([]int, size)
		// 处理本层的每个节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 根据方向决定填充位置
			if leftToRight {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
		leftToRight = !leftToRight // 下一层翻转方向
	}

	return res
}
