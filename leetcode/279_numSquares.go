package main

import (
	"fmt"
	"math"
)

func TestNumSquares() {
	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	candidates := make([]int, 0)
	for i := 1; i < n; i++ {
		if i*i > n {
			break
		}
		candidates = append(candidates, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, c := range candidates {
		for i := c; i < len(dp); i++ {
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}

	return dp[len(dp)-1]
}

func numSquares2(n int) int {
	op := make([]int, 0)
	for i := 1; i <= n; i++ {
		if i*i > n {
			break
		}
		op = append(op, i*i)
	}

	var dfs func(i, target int) bool
	dfs = func(i, target int) bool {
		if i <= 0 {
			return false
		}

		for j := 0; j < len(op); j++ {
			if op[j] > target {
				break
			}
			if op[j] == target {
				return true
			}
			if dfs(i-1, target-op[j]) {
				return true
			}
		}

		return false
	}

	for i := 1; i <= n; i++ {
		if dfs(i, n) {
			return i
		}
	}
	return -1
}
