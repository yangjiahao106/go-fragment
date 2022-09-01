package main

import "fmt"

func TestExist() {

	ans := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "DECSECBA")
	fmt.Println(ans)

}

func exist(board [][]byte, word string) bool {

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
			return false
		}

		board[i][j] = 0
		find := dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
		board[i][j] = word[k]
		return find
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
