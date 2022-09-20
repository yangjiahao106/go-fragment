package main

import "fmt"

func TestCandy() {
	fmt.Println(candy2([]int{0, 1}))
	fmt.Println(candy2([]int{0, 1, 2}))
	fmt.Println(candy2([]int{3,2,1, 2, 4, 3,4,5}))
}

func candy(ratings []int) int {
	dp := make([]int, len(ratings))
	dp[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			dp[i] = max(dp[i], dp[i+1]+1)
		}
	}
	//fmt.Println(dp)
	ans := 0
	for _, v := range dp {
		ans += v
	}

	return ans
}

// 扩展题  环形怎么处理
func candy2(ratings []int) int {
	dp := make([]int, len(ratings))
	dp[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			dp[i] = max(dp[i], dp[i+1]+1)
		}
	}
	fmt.Println(dp)

	// 两端进行比较
	if ratings[0] > ratings[len(ratings)-1] && dp[0] <= dp[len(dp)-1] {
		dp[0] = dp[len(dp)-1] + 1
		for i := 1; i < len(ratings); i++ {
			if ratings[i] > ratings[i-1] {
				dp[i] = dp[i-1] + 1
			}
		}
	}

	if ratings[0] < ratings[len(ratings)-1] && dp[0] >= dp[len(dp)-1] {
		dp[len(dp)-1] = dp[0] + 1
		for i := len(ratings) - 2; i >= 0; i-- {
			if ratings[i] > ratings[i+1] {
				dp[i] = max(dp[i], dp[i+1]+1)
			}
		}
	}

	fmt.Println( dp)
	ans := 0
	for _, v := range dp {
		ans += v
	}

	return ans
}
