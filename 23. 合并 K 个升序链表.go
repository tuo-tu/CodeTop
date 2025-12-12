package main

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
