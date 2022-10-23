package main

import "fmt"

func TestRemain() {
	fmt.Println(lastRemaining(10, 1))
	fmt.Println(lastRemaining2(10, 1))
}

func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}

	return (lastRemaining(n-1, m) + m) % n
}

func lastRemaining2(n int, m int) int {
	/*
	    https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solution/jian-zhi-offer-62-yuan-quan-zhong-zui-ho-dcow/
		f(n) = (f(n−1)+t)%n
	         = (f(n−1)+m%n)%n
		     = (f(n−1)+m)%n
	*/
	//
	x := 0
	for i := 2; i <= n; i++ {
		x = (x + m) % i
	}
	return x
}
