package main

// 核心思想：分组翻转 + 哨兵节点
// 使用 hair 作为哨兵节点永远指向头，dummy 标记当前组的前一个节点
// 对每 k 个节点为一组进行翻转：
//   1. 先找到当前组的尾节点 tail
//   2. 如果不足 k 个节点，直接返回（不翻转）
//   3. 记录下一组的开头，翻转当前组
//   4. 将翻转后的组连接到链表中
//   5. 更新 dummy 和 head，继续处理下一组
// 时间复杂度：O(n)，空间复杂度：O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	// hair 永远不会动,指向头
	hair := &ListNode{Next: head}
	// 一开始的dummy也指向头，但是后续会变
	dummy := hair
	// 需要找到每一段的 head、tail节点,每一组作为 for 的一个循环
	for head != nil {
		// 需要通过 dummy 来找到这一组的尾巴
		tail := dummy
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 如果这组元素不足，直接返回结果
			if tail == nil {
				return hair.Next
			}
		}
		// 记录下一组的开头
		nextHead := tail.Next
		// 翻转，返回新的头和尾
		newHead, newTail := reverse(head, tail)
		// 整合连接上
		dummy.Next = newHead
		newTail.Next = nextHead
		// 下一组开始
		// dummy等于上一组的尾节点
		dummy = newTail
		head = nextHead
	}
	return hair.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	// 记录本组的尾节点的下一个节点，从这里开始翻转
	pre := tail.Next
	cur := head
	for pre != tail {
		next := cur.Next
		// 当前节点指向前驱
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
