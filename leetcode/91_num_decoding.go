package main

import "fmt"

func testNumDecodings() {
	req := "0"
	ans := numDecodings_dp(req)
	fmt.Println(ans)

	ans = numDecodings(req)
	fmt.Println(ans)

}

/*
1

12
1 2

12 1
1 21
1 2 1

12 12
12 1 2
1 21 2
12 1 2
12 12
*/
func numDecodings(s string) int {
	if len(s) <= 1 || s[0] == '0' {
		return 0
	}

	var ans = 1
	singleEnd := 1 // 结尾为单独数字的数量
	for i := 1; i < len(s); i++ {
		tmp := 0
		if s[i] == '0' && s[i-1]-'0' >= 2 {
			return 0 // 0 不能单独存在，必须结合
		}
		if s[i] > '0' {
			tmp = ans // >0 可以单独存在
		}
		// 根据前一个数字判断是否可以结合
		if (s[i-1]-'0')*10+(s[i]-'0') <= 26 { // 可以结合
			tmp += singleEnd
			singleEnd = ans // singleEnd  + (ans_ - singleEnd)
		} else { // 无法结合， s[i] 全部单独存在
			singleEnd = ans
		}
		ans = tmp
	}

	return ans
}

func numDecodings_dp(s string) int {
	if len(s) <= 1 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		if s[i-1] >= '1' { // 可以单独存在
			dp[i] = dp[i-1]
		}
		v := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if v >= 10 && v <= 26 { // 可以结合。 dp[i-2] == singleEnd
			dp[i] += dp[i-2]
		}
	}

	return dp[len(dp)-1]
}
