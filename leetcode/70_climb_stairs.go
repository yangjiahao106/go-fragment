package main

import "fmt"

func TestClimbStairs() {
	ans := climbStairs7(45)
	fmt.Println(ans)

}

func climbStairs(n int) int {

	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return b
}

// 变种 ：不能爬到7及7的倍数
func climbStairs7(n int) int {

	a, b := 0, 1
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			a, b = b, 0
		} else {
			a, b = b, a+b
		}

	}
	return b
}
