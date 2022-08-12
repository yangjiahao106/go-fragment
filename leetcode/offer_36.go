package main

import "fmt"

// Definition for a Node.

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyListTest() {

	root := &Node {
		Val: 4,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 1,
			},
			Right: &Node{
				Val: 3,
			},
		},
		Right: &Node{
			Val: 5,
		},
	}

	treeToDoublyList(root)

	for i := 0; i < 10; i++ {
		fmt.Println(root.Val)
		root = root.Left
	}

}

func treeToDoublyList(root *Node) *Node {
	l, r := helper(root)
	l.Left = r
	r.Right = l
	return root
}

func helper(root *Node) (left *Node, right *Node) {
	if root == nil {
		return
	}
	left, right = root, root
	l1, r1 := helper(root.Left)
	l2, r2 := helper(root.Right)

	if r1 != nil {
		r1.Right = root
		root.Left = r1
		left = l1
	}
	if l2 != nil {
		root.Right = l2
		l2.Left = root
		right = r2
	}

	return
}
