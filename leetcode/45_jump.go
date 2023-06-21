package main

import (
	"fmt"
	"math"
)

func testJump() {
	fmt.Println(jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))

}

func jumpDP(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	for s, v := range nums {
		for i := s + v; i >= s; i-- {
			if i >= len(dp)-1 {
				return dp[s] + 1
			}
			if dp[i] > dp[s]+1 {
				dp[i] = dp[s] + 1
			} else {
				break
			}
		}
	}

	return dp[len(dp)-1]
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	step := 0
	i := 0

	for i+nums[i] < len(nums)-1 {
		farthest := 0
		to := i
		for j := i + 1; j <= i+nums[i]; j++ {
			if j+nums[j] > farthest { // 判断跳到哪里下次可以跳的最远
				to = j
				farthest = j + nums[j]
			}
		}
		i = to
		step += 1
	}
	return step + 1
}
