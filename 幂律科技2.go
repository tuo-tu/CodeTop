package main

import (
	"fmt"
)

// 创建几个Node对象，利用两个元素确定了他们的链接关系，是个树结构，根节点的父ID是“-1”，得出所有的路径。
// 比如A是根节点，子节点是A-1，A-1的子节点是A-1-1，以/A/A-1/A-1-1/的格式返回。
type node struct {
	ID  string
	PID string
}

func getAllPaths(nodes []node) []string {
	// 存储结果
	res := []string{}
	// 存储本节点和父节点
	nodeMap := map[string]string{}
	for _, node := range nodes {
		nodeMap[node.ID] = node.PID
	}

	// 递归处理当前节点, 从当前节点出发一直到根节点的路径
	var path func(id string) string
	path = func(id string) string {
		pid := nodeMap[id]
		if pid == "-1" {
			return id
		}
		// 否则继续往下遍历
		return path(pid) + "/" + id
	}

	// 遍历切片
	for _, node := range nodes {
		res = append(res, "/"+path(node.ID)+"/")
	}
	return res
}

func Test2Main() {
	nodes := []node{
		{"A", "-1"},
		{"A-1", "A"},
		{"A-2", "A"},
		{"A-3", "A"},
		{"A-2-1", "A-2"},
		{"A-2-2", "A-2"},
		{"A-2-3", "A-2"},
		{"A-2-4", "A-2"},
	}

	paths := getAllPaths(nodes)
	for _, path := range paths {
		fmt.Println(path)
	}
}
