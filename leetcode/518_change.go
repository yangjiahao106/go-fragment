package main

import "fmt"

func TestChane() {

	fmt.Println(change(12, []int{1, 2, 5}))
}

func change(amount int, coins []int) int {

	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i < len(dp); i++ {
			dp[i] += dp[i-coin]
		}
		//fmt.Println(dp)
	}

	return dp[len(dp)-1]
}
