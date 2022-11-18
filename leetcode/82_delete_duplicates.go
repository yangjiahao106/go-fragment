package main

import "fmt"

func TestDeleteDuplicates() {

	dummy := &ListNode{}
	cur := dummy
	for _, v := range []int{1,1, 2, 2, 2, 3, 4, 5} {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	ans := deleteDuplicates(dummy.Next)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head

	for fast != nil && fast.Next != nil {
		if fast.Val == fast.Next.Val {
			v := fast.Val
			for fast != nil && fast.Val == v {
				fast = fast.Next
			}
			slow.Next = fast

		} else {
			slow = fast
			fast = fast.Next
		}
	}

	return dummy.Next
}
