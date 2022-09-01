package main

import "fmt"

func TestFib() {
	fmt.Println(fib(5))
}

// 记忆化递归
func fib(n int) int {
	m := make(map[int]int)
	var helper func(n int) int
	helper = func(n int) int {
		if n <= 1 {
			return n
		}
		if v, ok := m[n]; ok {
			return v
		}
		ans := (helper(n-1) + helper(n-2)) % 1000000007
		m[n] = ans
		return ans
	}
	return helper(n)
}

// 动态规划
func fib2(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1

	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
