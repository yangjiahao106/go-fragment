package main

import (
	"fmt"
)

/*
		5
	3 			6
2	  4    	  5.5    8
	3.5 4.5     5.6
*/

func TestDeleteNode() {
	coder := Constructor()
	/*
				5
			3      6
		 2   4    X	 8

	*/
	root := coder.deserializeV3("5,3,6,2,4,X,7")
	fmt.Println(coder.serializeV3(root))

	root = deleteNode(root, 7)
	fmt.Println(coder.serializeV3(root))
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var helper func(root *TreeNode) *TreeNode
	helper = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == key {
			if root.Left == nil {
				return root.Right
			}
			if root.Right == nil {
				return root.Left
			}
			//if root.Left != nil && root.Right != nil {
			//  查找右子树的最小值, 取代 root
			var p *TreeNode
			cur := root.Right
			for cur.Left != nil {
				p = cur
				cur = cur.Left
			}
			if p != nil {
				p.Left = cur.Right
				cur.Right = root.Right
			}
			cur.Left = root.Left
			return cur

		} else if root.Val > key {
			root.Left = helper(root.Left)
		} else {
			root.Right = helper(root.Right)
		}
		return root
	}

	return helper(root)
}
