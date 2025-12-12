package main

// DFS + 递归
// 深度优先搜索（DFS）的典型特征：一条路走到黑，再回头走别的路。
func levelOrder(root *TreeNode) [][]int {
	// result用于存储结果，二维切片的每个一维索引代表一层
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 如果当前层不存在，先新建一层
		// len(res)表示已经存放了几层
		// 比如目前已经存放了根节点（0层），此时len(res) == 1，但是第1层还不存在，因此要创建[]int{}
		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)
		// 递归处理左右子树
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return res
}
