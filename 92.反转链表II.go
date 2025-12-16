package main

// 一次遍历「穿针引线」反转链表（头插法）
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 设置哨兵节点
	dummy := &ListNode{Next: head}
	pre := dummy
	// 获取left的前一个pre节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	
	// 切记不要这样写，因为指针的指向都变了
	// next.Next = cur
	// cur = next
	// 开始翻转，头插法
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next     // 1.选取要头插的节点
		cur.Next = next.Next // 2.从原位置删掉 next
		next.Next = pre.Next // 3.将 next 指向翻转部分的头（永远是 pre.next）,千万不能写成 cur
		pre.Next = next      // 4.pre 指向 next，完成插入。千万不能写成 cur = next
	}
	return dummy.Next
}
