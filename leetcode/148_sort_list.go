package main

import (
	"fmt"
)

func TestSortList() {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range []int{9,8} {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	ans := sortList2(dummy.Next)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

// quick sort
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	head = head.Next
	// 将链表拆分成大小两条链表
	low := &ListNode{}
	lowCur := low
	high := &ListNode{Next: head}
	cur := high
	for cur.Next != nil {
		if cur.Next.Val <= p.Val {
			lowCur.Next = cur.Next
			lowCur = lowCur.Next
			cur.Next = cur.Next.Next
			lowCur.Next = nil
		} else {
			cur = cur.Next
		}
	}

	l := sortList2(low.Next)
	r := sortList2(high.Next)
	if l == nil {
		p.Next = r
		return p
	}
	// 连接链表
	cur = l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = p
	p.Next = r
	return l
}

// merge sort
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)

	return sortListMerge(left, right)
}

func sortListMerge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}
	return dummy.Next
}


func sortList3(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}

	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = sortListMerge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}
