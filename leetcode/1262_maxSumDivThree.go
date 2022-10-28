package main

import "fmt"

func TestMaxSumDivThree() {

	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))

}

func maxSumDivThree(nums []int) int {
	dp := make([]int, 3)

	for _, v := range nums {
		a := dp[0] + v
		b := dp[1] + v
		c := dp[2] + v

		dp[a%3] = max(dp[a%3], a)
		dp[b%3] = max(dp[b%3], b)
		dp[c%3] = max(dp[c%3], c)

	}

	return dp[0]
}
