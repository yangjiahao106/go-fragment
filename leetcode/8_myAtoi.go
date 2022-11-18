package main

import (
	"fmt"
	"math"
)

func TestMyAtoi() {
	//fmt.Println(math.MaxInt32, math.MinInt32)

	fmt.Println(myAtoi("  +00234  sldkfwe"))
}

func myAtoi(s string) int {
	ans := 0
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	flag := 1
	if i < len(s) {
		if s[i] == '+' {
			i++
		} else if s[i] == '-' {
			flag = -1
			i++
		}
	}

	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			x := int(s[i] - '0')
			if ans > (math.MaxInt32-x)/10 {
				if flag == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
			ans = ans*10 + x
		} else {
			break
		}
	}

	return ans * flag
}
