package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for i := coin; i < len(dp); i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[len(dp)-1] == math.MaxInt32 {
		return -1
	}

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
