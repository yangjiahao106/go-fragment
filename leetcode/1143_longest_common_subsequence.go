package main

import "fmt"

func TestLongestCommonSubsequence() {
	ans := longestCommonSubsequence("sdgwesadgwe", "sdgweasdgsdfwe")
	fmt.Println(ans)

}

func longestCommonSubsequence(text1 string, text2 string) int {

	dp := make([][]int, len(text1)+1)
	for k := range dp {
		dp[k] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	length := dp[len(dp)-1][len(dp[0])-1]

	// 输出 路径
	for _, v := range dp {
		fmt.Println(v)
	}
	path := make([]byte, length)
	r := len(dp[0]) - 1
	cur := length - 1
	for i := len(dp) - 1; i > 0; i-- {
		for j := r; j > 0; j-- {
			if dp[i][j] > dp[i-1][j-1] && dp[i-1][j-1] == dp[i-1][j] && dp[i-1][j-1] == dp[i][j-1] {
				path[cur] = text1[i-1]
				cur--
				r = j - 1 // 设置有边界 避免重复
				break     // break 防止重复
			}
		}
	}
	fmt.Println(string(path))

	return length
}
