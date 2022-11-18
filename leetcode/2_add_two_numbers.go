package main

import "fmt"

func TestAddTwoNumbers() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	ans := addTwoNumbers(l1, l2)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	remain := 0
	cur := dummy
	for l1 != nil || l2 != nil || remain > 0 { // 注意 remain > 0
		if l1 != nil {
			remain += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remain += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: remain % 10}
		remain = remain / 10
		cur = cur.Next
	}

	return dummy.Next
}
