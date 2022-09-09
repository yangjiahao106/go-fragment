package main

import "fmt"

/*
class Solution:
    def backToOrigin(self,n):
        #点的个数为10
        length = 10
        dp = [[0 for i in range(length)] for j in range(n+1)]
        dp[0][0] = 1
        for i in range(1,n+1):
            for j in range(length):
                #dp[i][j]表示从0出发，走i步到j的方案数
                dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
        return dp[n][0]

dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length] 


*/

func TestBackToOrigin() {
	backToOrigin2(10)
}

func backToOrigin(n int) int {
	length := 10

	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, length)
	}
	dp[0][0] = 1

	fmt.Println(dp[0])

	for i := 1; i <= n; i++ {
		for j := 0; j < length; j++ {
			//dp[i][j]表示从0出发，走i步到j的方案数
			dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
		}
		fmt.Println(dp[i])
	}
	return dp[n][0]
}

func backToOrigin2(n int) int {
	length := 10
	dp := make([]int, n)
	dp[0] = 1
	fmt.Println(dp)

	for i := 1; i <= n; i++ {
		temp := make([]int, n)
		for j := 0; j < length; j++ {
			temp[j] = dp[(j+length-1)%length] + dp[(j+1)%length]
		}
		dp = temp
		fmt.Println(dp)
	}
	return dp[0]
}
