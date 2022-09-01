package main

import "fmt"

func testRotateRight(){
	ans := rotateRight(nil, 0)
	fmt.Println(ans)

	// github.com  yangjiaaho.github.com go to the hell profit exist
	// function button logic keyboard is very hard to use
	// delete pause b home end page up page down eeww er fuhfh keli de
	// rotate  coin change ddigits head cur.Next skdeskks
	// permute uni que sum conbination sum area max min su array calculate zhuang bi ne ba ni
	//
	// permute uniqu number list node list node k heasd is not the head hahaa cur.Next head length head

	// rotate right head length
	// rotate the list right users edz go go setup lib exec
	// mo he hao le ma ma
	// rotate righe lsk jdielkjfjjjj left right go go go exec command library
}

func rotateRight(head *ListNode, k int) *ListNode {

	length := 0
	cur := head
	tail := head
	for cur != nil {
		tail = cur
		cur = cur.Next
		length += 1
	}

	k = k % length
	if k == 0 {
		return head
	}
	pre := head

	for i := 0; i < length-k-1; i++ {
		pre = pre.Next
	}

	newHead := pre.Next
	pre.Next = nil
	tail.Next = head
	return newHead
}
