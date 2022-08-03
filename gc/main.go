package main

import "fmt"

func main() {
	s := make([]int, 1000)
	_ = s

	s2 := make([]int, 10000)
	_ = s2

	s3 := make([]int, 100)
	_ = s3

	foo(s3)
}

func foo(ss []int) {
	fmt.Println(ss)
}
