package main

func BuildTreeTest() {

	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

}

func buildTree(preorder []int, inorder []int) *TreeNode {

	var dfs func(inorder []int) *TreeNode
	dfs = func(inorder []int) *TreeNode {
		if len(inorder) == 0 {
			return nil
		}
		v := preorder[0]
		preorder = preorder[1:]

		i := 0
		for ; i < len(inorder); i++ {
			if inorder[i] == v {
				break
			}
		}

		node := &TreeNode{Val: v}
		node.Left = dfs(inorder[:i])
		node.Right = dfs(inorder[i+1:])

		return node
	}

	return dfs(inorder)
}
