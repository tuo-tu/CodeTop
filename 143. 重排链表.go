package main

// 快慢指针找中点 + 反转后半段 + 交叉合并
func reorderList1(head *ListNode) {
	// 特殊情况
	if head == nil {
		return
	}

	// 1.找到中间节点 slow
	slow, fast := head, head
	// 注意不要写成 slow.Next != nil ，统一移动fast
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2.翻转后半部分,pre 就是新的头节点
	var pre *ListNode
	cur := slow.Next
	// 注意要断开前后部分的联系
	slow.Next = nil
	// 开始翻转后半部分链表
	for cur != nil {
		next := cur.Next
		// 不要写成 next.Next = pre
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 3.交叉合并，注意前半部分长度一定是 >= 后半部分长度
	head1, head2 := head, pre
	// 大错特错的写法：
	// 执行 head1 = head1.Next 之后 head1 恰好变成 head2，
	// 随后的 head2.Next = head1 就把 head2.Next 指向自己，立刻产生自环，
	// 导致循环不终止或链表结构被破坏。
	//for head2 != nil {
	//	head1.Next = head2
	//	head1 = head1.Next
	//	head2.Next = head1
	//	head2 = head2.Next
	//}

	for head2 != nil {
		t1, t2 := head1.Next, head2.Next
		head1.Next = head2
		// 不要写成 t2.Next = t1
		head2.Next = t1
		// 不要写成 t1, t2 = t2, t1.Next
		head1, head2 = t1, t2
	}
}

// 数组辅助
func reorderList(head *ListNode) {
	// 特殊情况
	if head == nil {
		return
	}

	nodes := []*ListNode{}
	// 存储链表的每个节点
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}

	// 双指针遍历切片
	left, right := 0, len(nodes)-1
	for left < right {
		nodes[left].Next = nodes[right]
		left++
		// 避免重复连接导致成环
		if left == right {
			break
		}
		nodes[right].Next = nodes[left]
		right--
	}
	// 把链表尾节点的 Next 指针置为 nil
	nodes[right].Next = nil
}
