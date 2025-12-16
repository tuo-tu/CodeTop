package main

// 核心思想：分治合并（两两合并）
// 将 k 个链表两两配对合并，重复这一过程直到只剩下一个链表
// 第一轮：合并 lists[0] 和 lists[1]，lists[2] 和 lists[3]，...
// 第二轮：合并前一轮的结果，每次合并的间隔翻倍
// 时间复杂度：O(kn log k)，其中 k 是链表数量，n 是每个链表的平均长度
// 空间复杂度：O(1)（不考虑递归栈空间）
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// 一共多少个链表
	n := len(lists)
	for interval := 1; interval < n; interval *= 2 {
		for i := 0; i+interval < n; i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
	}
	// 最后只剩下一个链表，返回这个链表
	return lists[0]
}

//func mergeTwoLists(l1, l2 *ListNode) *ListNode {
//	dummy := &ListNode{}
//	cur := dummy
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			cur.Next = l1
//			l1 = l1.Next
//		} else {
//			cur.Next = l2
//			l2 = l2.Next
//		}
//		cur = cur.Next
//	}
//	if l1 != nil {
//		cur.Next = l1
//	} else {
//		cur.Next = l2
//	}
//	return dummy.Next
//}
