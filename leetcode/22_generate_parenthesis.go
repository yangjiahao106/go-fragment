package main

import "fmt"

func TestGenerateParentthesis() {

	// (((   )))
	/*

	 */

	ans := generateParenthesis(4)
	fmt.Println(ans)
}

func generateParenthesis(n int) []string {

	ans := make([]string, 0)
	path := make([]byte, 0, n*2)

	var dfs func(path []byte, left, right int)
	dfs = func(path []byte, left, right int) {
		if len(path) == n*2 {
			ans = append(ans, string(path))
		}
		if left > 0 {
			dfs(append(path, '('), left-1, right)
		}

		if left < right {
			dfs(append(path, ')'), left, right-1)
		}
	}

	dfs(path, n, n)

	return ans
}
