package main

// 哨兵节点不动，cur节点动
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 初始化哨兵节点,不需要增加 cur1，cur2
	dummy := &ListNode{}
	// cur 和 dummy 指向同一个节点,但是cur会往前移动，dummy不会
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			// 注意链表要往前滑动
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		// cur 也要往前滑动
		cur = cur.Next
	}

	// 还要拼接剩下的链表
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}
