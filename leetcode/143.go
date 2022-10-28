package main

import "fmt"

func TestReorderList() {
	// 1，,2，,3，,4，,5
	dummy := &ListNode{}
	cur := dummy
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	reorderList(dummy.Next)

	for dummy.Next != nil {
		fmt.Println(dummy.Next.Val)
		dummy.Next = dummy.Next.Next
	}

}

// 寻找链表中点
func reorderList(head *ListNode) {
	// find middle
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse
	cur := slow.Next
	slow.Next = nil
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head2 := pre

	// connect
	cur = &ListNode{}
	for head != nil || head2 != nil {
		if head != nil {
			cur.Next = head
			head = head.Next
			cur = cur.Next
		}
		if head2 != nil {
			cur.Next = head2
			head2 = head2.Next
			cur = cur.Next
		}
	}
}

func reorderList2(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	middle := slow.Next
	slow.Next = nil

	var pre *ListNode
	for middle != nil {
		next := middle.Next
		middle.Next = pre
		pre = middle
		middle = next
	}
	// 1, 2, 3
	// 5, 4,

	head1 := head
	head2 := pre
	for head1 != nil && head2 != nil {
		next1, next2 := head1.Next, head2.Next
		head1.Next = head2
		head2.Next = next1
		head1 = next1
		head2 = next2
	}
}
