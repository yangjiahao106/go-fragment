package main

import "fmt"

func TestCanJump() {

	fmt.Println(canJump([]int{3, 2, 1, 1, 4}))
}

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for k, v := range nums {
		if dp[k] {
			for i := k + v; i > k; i-- {
				if i >= len(dp)-1 {
					return true
				}
				if dp[i] {
					break
				}
				dp[i] = true
			}
		}
	}
	return dp[len(dp)-1]
}

func canJump2(nums []int) bool {
	longest := 0

	for k, v := range nums {
		if longest >= k && k+v > longest {
			longest = k + v
		}
	}

	return longest >= len(nums)-1
}
