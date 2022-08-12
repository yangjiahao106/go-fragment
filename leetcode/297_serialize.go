package main

import (
	"strconv"
	"strings"
)

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{

	}
}

// V1 DFS
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var helper func(root *TreeNode)
	b := strings.Builder{}
	helper = func(root *TreeNode) {
		if root == nil {
			b.WriteString("X,")
			return
		}

		b.WriteString(strconv.Itoa(root.Val))
		b.WriteByte(',')

		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	return b.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nums := strings.Split(data, ",")

	var helper func() *TreeNode
	helper = func() *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		n := nums[0]
		nums = nums[1:]
		if n == "X" {
			return nil
		}
		v, _ := strconv.Atoi(n)
		return &TreeNode{
			Val:   v,
			Left:  helper(),
			Right: helper(),
		}
	}
	return helper()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// V2

// Serializes a tree to a single string.
func (this *Codec) serializeV2(root *TreeNode) string {
	var helper func(root *TreeNode)
	b := strings.Builder{}

	helper = func(root *TreeNode) {
		if root == nil {
			b.WriteString("X")
			return
		}

		b.WriteByte('(')
		helper(root.Left)
		b.WriteByte(')')

		b.WriteString(strconv.Itoa(root.Val))

		b.WriteByte('(')
		helper(root.Right)
		b.WriteByte(')')
	}

	helper(root)
	return b.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeV2(data string) *TreeNode {

	var helper func() *TreeNode
	helper = func() *TreeNode {

		if data[0] == 'X' {
			data = data[1:]
			return nil
		}

		data = data[1:]
		left := helper()
		data = data[1:]

		i := 0
		for data[i] == '-' || data[i] >= '0' && data[i] <= '9' {
			i++
		}

		v, _ := strconv.Atoi(data[:i])
		data = data[i:]
		data = data[1:]
		right := helper()
		data = data[1:]
		return &TreeNode{
			Val:   v,
			Left:  left,
			Right: right,
		}
	}
	return helper()
}

// V3 BFS

// Serializes a tree to a single string.
func (this *Codec) serializeV3(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := make([]string, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			ans = append(ans, "X")
		} else {
			ans = append(ans, strconv.Itoa(n.Val))
			queue = append(queue, n.Left)
			queue = append(queue, n.Right)
		}
	}
	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeV3(data string) *TreeNode {
	if len(data) == 0{
		return nil
	}
	nums := strings.Split(data, ",")
	v, _ := strconv.Atoi(nums[0])
	nums = nums[1:]

	root := &TreeNode{
		Val: v,
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(nums) > 0 {
		node := queue[0]
		queue = queue[1:]

		l := nums[0]
		nums = nums[1:]
		if l != "X" {
			v, _ := strconv.Atoi(l)
			left := &TreeNode{Val: v}
			node.Left = left
			queue = append(queue, left)
		}
		r := nums[0]
		nums = nums[1:]
		if r != "X" {
			v, _ := strconv.Atoi(r)
			right := &TreeNode{Val: v}
			node.Right = right
			queue = append(queue, right)
		}
	}

	return root
}
