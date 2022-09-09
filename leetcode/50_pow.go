package main

import (
	"fmt"
	"math"
)

func TestMyPow() {
	fmt.Println(math.Pow(2.0001, 33))
	fmt.Println(myPow(2.0001, 33))
	fmt.Println(myPow2(2.0001, 33))
}

func myPow(x float64, n int) float64 {
	dp := make(map[int]float64)
	var helper func(x float64, n int) float64
	helper = func(x float64, n int) float64 {
		if ans, ok := dp[n]; ok {
			return ans
		}
		var ans float64
		switch n {
		case 0:
			ans = 1
		case 1:
			ans = x
		default:
			ans = helper(x, n/2) * helper(x, n-n/2)
		}
		dp[n] = ans
		return ans
	}

	if n < 0 {
		return 1 / helper(x, -n)
	}
	return helper(x, n)
}

func myPow2(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPowHelper(x, -n)
	}
	return myPowHelper(x, n)
}

func myPowHelper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	return myPowHelper(x, n/2) * myPowHelper(x, n-n/2)
}


func myPow3(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, N int) float64 {
	ans := 1.0
	// 贡献的初始值为 x
	x_contribute := x
	// 在对 N 进行二进制拆分的同时计算答案
	for N > 0 {
		if N % 2 == 1 {
			// 如果 N 二进制表示的最低位为 1，那么需要计入贡献
			ans *= x_contribute
		}
		// 将贡献不断地平方
		x_contribute *= x_contribute
		// 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
		N /= 2
	}
	return ans
}

