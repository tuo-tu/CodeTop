package main

// 哈希表
func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]bool{}
	cur := head
	for cur != nil {
		if visited[cur] {
			return true
		}
		visited[cur] = true
		cur = cur.Next
	}
	return false
}

// 快慢指针
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	// 如果 fast.Next != nil，那么就代表fast不为空，不是也把 fast != nil包含进去了吗？fast != nil是否多余？
	// 如果 能安全地访问到 fast.Next，那它必然不是 nil。
	// 但是 问题在于 Go 的求值顺序 ——在判断条件里，编译器必须先去取出 fast.Next 的值，才能知道它是不是 nil。
	// 如果这时 fast == nil，取 fast.Next 就会直接触发 panic，而不会等你去判断结果。
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步

		if slow == fast { // 相遇说明有环
			return true
		}
	}
	return false
}
