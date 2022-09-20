package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		k--
		if k == 0 {
			ans = root.Val
			return
		}
		dfs(root.Right)
	}

	dfs(root)
	return ans
}

// 非递归 中序遍历
func kthSmallest2(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}

// 频繁查询一个树， 将每个树的左右节点树记录下来，用二分法查找
type KthSmallest struct {
	root *TreeNode
	m    map[*TreeNode]int
}

func (ks *KthSmallest) KthSmallest(k int) int {
	root := ks.root

	for {
		if ks.m[root.Left] < k-1 {
			k -= ks.m[root.Left] + 1
			root = root.Right
		} else if ks.m[root.Left] > k-1 {
			root = root.Left
		} else {
			return root.Val
		}

	}
	//return -1
}

func (ks *KthSmallest) Count() {
	ks.helper(ks.root)
}

func (ks *KthSmallest) helper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := ks.helper(node.Left) + ks.helper(node.Right) + 1
	ks.m[node] = count

	return count
}

func TestKthSmallest() {
	//        6
	//    2       8
	//  1  3    7   9
	//  	4
	//  	 5

	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3,
				Right: &TreeNode{
					Val:   4,
					Right: &TreeNode{Val: 5},
				},
			},
		},

		Right: &TreeNode{Val: 8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}

	ks := KthSmallest{
		root: root,
		m:    map[*TreeNode]int{},
	}
	ks.Count()
	fmt.Println(ks.m)
	for i := 1; i <= 9; i++ {
		fmt.Println(ks.KthSmallest(i))
	}
}
