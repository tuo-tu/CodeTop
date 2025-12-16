package main

func reverseList(head *ListNode) *ListNode {
	// 不能写成 pre := &ListNode{}，因为这里不需要哨兵节点，如果这样写，最后尾巴会多一个0值节点
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // 先记录下一个节点的位置
		cur.Next = pre   // 反转当前节点,指向前驱
		pre = cur        // pre 先前进
		cur = next       // cur 再前进
	}
	return pre
}
