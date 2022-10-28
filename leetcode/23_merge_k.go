package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestMergeKLists() {

	list := make([]*ListNode, 0)
	list = append(list, &ListNode{Val: 1})
	list = append(list, &ListNode{Val: 1})
	list = append(list, &ListNode{Val: 1})

	head := mergeKLists(list)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

type PriorityQueue []*ListNode

func (qPtr *PriorityQueue) Push(node *ListNode) {
	q := *qPtr
	q = append(q, node)

	idx := len(q) - 1
	for idx > 1 && q[idx].Val < q[idx/2].Val {
		q[idx], q[idx/2] = q[idx/2], q[idx]
		idx = idx / 2
	}

	*qPtr = q
}

func (qPtr *PriorityQueue) Pop() *ListNode {
	q := *qPtr

	ret := q[1]
	q[1] = q[len(q)-1]

	q = q[:len(q)-1]
	idx := 1
	for {
		next := idx
		if idx*2 < len(q) && q[idx].Val > q[idx*2].Val {
			next = idx * 2
		}
		if idx*2+1 < len(q) && q[next].Val > q[idx*2+1].Val {
			next = idx*2 + 1
		}
		if next == idx {
			break
		}
		q[idx], q[next] = q[next], q[idx]
		idx = next
	}

	*qPtr = q
	return ret
}

// 优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	q := PriorityQueue{&ListNode{}}

	for _, v := range lists {
		if v != nil {
			q.Push(v)
		}
	}
	fmt.Println(q)

	dummy := &ListNode{}
	pre := dummy

	for len(q) > 1 {
		node := q.Pop()
		fmt.Println(q)

		if node.Next != nil {
			q.Push(node.Next)
		}
		pre.Next = node
		pre = pre.Next
	}
	return dummy.Next
}

// 优先队列
func mergeKLists3(lists []*ListNode) *ListNode {
	q := make([]*ListNode, 1)

	for _, v := range lists {
		q = priorityQueuePush(q, v)
	}

	dummy := &ListNode{}
	pre := dummy

	for len(q) > 1 {
		node := q[1]
		q = priorityQueuePop(q)
		if node.Next != nil {
			q = priorityQueuePush(q, node.Next)
		}

		pre.Next = node
		pre = pre.Next

	}

	return dummy.Next
}

func priorityQueuePush(q []*ListNode, node *ListNode) []*ListNode {
	q = append(q, node)
	idx := len(q) - 1

	for idx > 1 && q[idx].Val < q[idx/2].Val {
		q[idx], q[idx/2] = q[idx/2], q[idx]
		idx = idx / 2
	}

	return q
}

func priorityQueuePop(q []*ListNode) []*ListNode {
	q[1] = q[len(q)-1]
	q = q[:len(q)-1]

	idx := 1
	for {
		next := idx
		if idx*2 < len(q) && q[idx].Val > q[idx*2].Val {
			next = idx * 2
		}

		if idx*2+1 < len(q) && q[next].Val > q[idx*2+1].Val {
			next = idx*2 + 1
		}

		if next == idx {
			break
		}

		q[idx], q[next] = q[next], q[idx]
		idx = next
	}

	return q
}

// 二分法
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	m := len(lists) / 2
	left := mergeKLists(lists[:m])
	right := mergeKLists(lists[m:])

	dummy := &ListNode{}
	pre := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			pre = left
			left = left.Next
		} else {
			pre = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left != nil {
		pre.Next = left
	}
	if right != nil {
		pre.Next = right
	}

	return dummy.Next
}
