package main

import (
	"fmt"
	"math/rand"
	"time"
)

func TestRand10() {
	m := map[int]int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100000; i++ {
		m[rand10_3()] += 1
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(i, m[i])
	}
}

func rand7() int {
	return 1 + rand.Intn(7)
}

func rand6() int {
	for {
		v := rand7()
		if v <= 6 {
			return v
		}
	}
}

func rand5() int {
	for {
		v := rand7()
		if v <= 5 {
			return v
		}
	}
}

func rand10() int {
	if rand6()%2 == 0 {
		return rand5()
	} else {
		return rand5() + 5
	}
}

// 解法2
func rand10_() int {
	// https://pic.leetcode-cn.com/1630776258-UNMORj-%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20210905012406.jpg
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col

		//fmt.Println(row, col, idx)

		if idx >= 10 {
			return idx%10 + 1
		}
	}
}

func rand10_3() int {
	return (rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7())%10 + 1

}
