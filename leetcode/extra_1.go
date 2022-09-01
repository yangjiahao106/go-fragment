package main

import "fmt"

// 附加 1
// 奇偶链表 拆分合并

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestSortOddEvenList() {
	//createList
	vals := []int{1, 8, 2, 7, 3, 6, 4, 5, 5, 4, 6, 3}

	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	ans := sortOddEvenList(dummy.Next)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func sortOddEvenList(head *ListNode) *ListNode {
	// partition
	evenHead := head.Next
	odd, even := head, evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = even.Next

		even.Next = odd.Next
		even = odd.Next
	}
	odd.Next = nil

	// reverse
	var pre *ListNode
	cur := evenHead
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	evenHead = pre

	// merge
	odd, even = head, evenHead
	dummy := &ListNode{}
	cur = dummy
	for odd != nil && even != nil {
		if odd.Val <= even.Val {
			cur.Next = odd
			odd = odd.Next
		} else {
			cur.Next = even
			even = even.Next
		}
		cur = cur.Next
	}
	if odd != nil {
		cur.Next = odd
	}
	if even != nil {
		cur.Next = even
	}

	return dummy.Next
}
