package main

import "fmt"

func TestReverseKGroup() {
	dummy := &ListNode{}
	pre := dummy
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range nums {
		pre.Next = &ListNode{Val: v}
		pre = pre.Next
	}

	head := reverseKGroup(dummy.Next, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	pre, cur := dummy, head

	for cur != nil {
		left, right := cur, cur
		for i := 1; i < k && right != nil; i++ {
			right = right.Next
		}
		if right == nil {
			pre.Next = left
			break // 不足k个跳出
		}

		cur = right.Next
		reverseKGroupHelper(left, right)
		pre.Next = right // 连上反转后的头节点
		pre = left
	}

	return dummy.Next
}

// reverse
func reverseKGroupHelper(head, tail *ListNode) {
	cur := head
	var pre *ListNode
	for {
		next := cur.Next
		cur.Next = pre
		if cur == tail {
			return
		}
		pre, cur = cur, next
	}
}
