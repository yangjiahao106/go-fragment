package main

import "fmt"

func TestOddEvenList() {
	dummy := &ListNode{}
	cur := dummy
	for i := 0; i <= 44; i++ {
		node := &ListNode{Val: i}
		cur.Next = node
		cur = cur.Next
	}

	ans := oddEvenList(dummy.Next)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}

/*
给定单链表的头节点head，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
第一个节点的索引被认为是 奇数 ， 第二个节点的索引为偶数 ，以此类推。
请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
你必须O(1)的额外空间复杂度和O(n)的时间复杂度下解决这个问题。
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head2 := head.Next

	pre1, pre2 := head, head2
	cur := head2.Next
	for cur != nil && cur.Next != nil {
		pre1.Next, pre2.Next = cur, cur.Next
		cur = cur.Next.Next
		pre1, pre2 = pre1.Next, pre2.Next
	}

	if cur != nil {
		pre1.Next, pre2.Next = cur, nil // pre2.Next = nil  防止死循环
		pre1 = pre1.Next
	}
	pre1.Next = head2

	return head
}
