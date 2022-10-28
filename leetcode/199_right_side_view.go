package main

func rightSideView(root *TreeNode) []int {

	ans := make([]int, 0)
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, d int) {
		if root == nil {
			return
		}
		if d == len(ans) {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, d+1)
		dfs(root.Left, d+1)
	}
	dfs(root, 0)
	return ans
}
