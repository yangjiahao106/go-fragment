package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s))
	for i := 0; i < len(dp); i++ {
		if i == 0 || dp[i-1] {
			for _, word := range wordDict {
				if i+len(word) <= len(dp) && s[i:i+len(word)] == word {
					dp[i+len(word)-1] = true
				}
			}
		}
	}

	return dp[len(dp)-1]
}

func wordBreakTest() {
	fmt.Println(wordBreak("pen", []string{"apple", "pen"}))
}
