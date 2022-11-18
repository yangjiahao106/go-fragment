package main

import (
	"fmt"
	"time"
)

func main() {
	TestOneProcessor()
	//rearrangement2()
	//readWrite()
}

func readWrite() {

	var i int64 = 0
	nums := []int64{
		1111111111111111111,
		2222222222222222222,
		3333333333333333333,
		4444444444444444444,
		5555555555555555555,
		6666666666666666666,
		7777777777777777777,
		8888888888888888888,

	}

	m := make(map[int64]bool)
	for _, v := range nums {
		m[v] = true
	}

	go func() {
		for {
			for v := range m {
				i = v
				//fmt.Println("set",v)
			}
		}
	}()

	go func() {
		for {
			for v := range m {
				i = v
				fmt.Println("set",v)
			}
		}
	}()

	go func() {
		for j := 0; ; j++ {
			fmt.Println("read",j,i)
			if !m[i]{
				panic(i)
			}
		}
	}()

	time.Sleep(time.Second/10)
}
