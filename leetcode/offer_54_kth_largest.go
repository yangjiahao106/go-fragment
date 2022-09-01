package main

func kthLargest(root *TreeNode, k int) int {
	count := 0
	ans := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil{
			return
		}
		dfs(root.Right)

		count += 1
		if count == k {
			ans = root.Val
			return
		}
		dfs(root.Left)

	}

	dfs(root)

	return ans
}
