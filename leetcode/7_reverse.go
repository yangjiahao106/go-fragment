package main

import (
	"fmt"
	"math"
)

func TestReverse7() {
	fmt.Println(reverse7(-1234))
}

func reverse7(x int) int {
	ans := 0
	for x != 0 {
		n := x % 10
		x = x / 10
		if ans > (math.MaxInt32-n)/10 {
			return 0
		}

		ans *= 10
		ans += n
	}
	return ans
}
