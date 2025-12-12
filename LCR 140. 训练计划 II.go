package main

//func trainingPlanII(head *ListNode, k int) *ListNode {
//	if head == nil /*|| k <= 0*/ {
//		return nil
//	}
//
//	fast, slow := head, head
//
//	// 让 fast 先走 k 步，这样slow和fast之间就隔了 k 个节点
//	for i := 0; i < k; i++ {
//		// 如果链表长度 < k，提前结束
//		if fast == nil {
//			return nil
//		}
//		fast = fast.Next
//	}
//
//	// 快慢同步走
//	for fast != nil {
//		fast = fast.Next
//		slow = slow.Next
//	}
//
//	return slow
//}
//
//// 构建链表
//func buildList(nums []int) *ListNode {
//	if len(nums) == 0 {
//		return nil
//	}
//	head := &ListNode{Val: nums[0]}
//	cur := head
//	for _, v := range nums[1:] {
//		cur.Next = &ListNode{Val: v}
//		cur = cur.Next
//	}
//	return head
//}
//
//// 打印链表
//func printList(head *ListNode) {
//	for head != nil {
//		fmt.Printf("%d ", head.Val)
//		head = head.Next
//	}
//	fmt.Println()
//}
//
//func main() {
//	// 示例链表: 1 -> 2 -> 3 -> 4 -> 5
//	head := buildList([]int{1, 2, 3, 4, 5})
//
//	/*fmt.Println("正常情况：")
//	res := trainingPlanII(head, 2) // 倒数第 2 个节点
//	printList(res)                 // 4 5
//
//	fmt.Println("边界测试：")
//	// 1. 空链表
//	fmt.Println(trainingPlanII(nil, 1)) // nil
//
//	// 2. k = 0
//	fmt.Println(trainingPlanII(head, 0)) // nil
//
//	// 3. k > 长度
//	fmt.Println(trainingPlanII(head, 10)) // nil
//
//	// 4. k = 1（尾节点）
//	printList(trainingPlanII(head, 1)) // 5
//
//	// 5. k = len(list)（头节点）
//	printList(trainingPlanII(head, 5))  // 1 2 3 4 5*/
//	printList(trainingPlanII(head, -2)) // 1 2 3 4 5
//}


func trainingPlanII(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head

	// 让 fast 先走 k 步
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	// 快慢同步走
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

// 构造链表：1 -> 2 -> 3 -> 4 -> 5
// func buildList() *ListNode {
// 	head := &ListNode{Val: 1}
// 	cur := head
// 	for i := 2; i <= 5; i++ {
// 		cur.Next = &ListNode{Val: i}
// 		cur = cur.Next
// 	}
// 	return head
// }

// func main() {
// 	head := buildList()

// 	// 正常情况 k=2，应该返回倒数第2个节点 (4)
// 	res1 := trainingPlanII(head, 2)
// 	fmt.Println("k=2, result:", res1) // 输出 4

// 	// 特殊情况 k=-1，应该返回最后一个节点 (5)
// 	res2 := trainingPlanII(head, -2)
// 	fmt.Println("k=-1, result:", res2) // 输出 5
// }
