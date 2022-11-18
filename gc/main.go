package main

import (
	"fmt"
	"runtime"
)

/*

查看内存分配情况

go build -gcflags "-m -l" main.go

*/

type ss struct {
	s string
}

func main() {
	runtime.Gosched()
	ip := new(int)   // new 不一定发生内存逃逸 ，可能分配在堆上也可能分配在栈上。
	sp := new([]int) // new 不一定发生内存逃逸 ，可能分配在堆上也可能分配在栈上。
	ssp := new(ss)   // new 不一定发生内存逃逸 ，可能分配在堆上也可能分配在栈上。

	_ = sp
	_ = ip
	_ = ssp
	s := make([]int, 1000)
	_ = s

	s2 := make([]int, 10000)
	_ = s2

	s3 := make([]int, 100)
	foo(s3)

	var i = 1
	var j = 1
	var k = 1
	receiveInterface(k)
	_ = i
	addInt(&i) // 不会导致逃逸, 向下传递指针未必会逃逸， 向外返回指针会发生逃逸
	is := make([]*int, 1)
	is[0] = &i     // 会导致逃逸
	fmt.Println(i) // print 内部会返回指针，会导致逃逸。

	pj := &j
	fmt.Println(&i)
	print(pj)

	s4 := [100000]int{}
	s5 := [100000]int{1, 2, 3, 4, 5}

	addArray(s4)
	printArray2(&s5)
}

func addInt(i *int) {
	*i += 1
}

func foo(ss []int) {
	fmt.Println(ss) // print 传递interface 会导致 ss 和 s3 逃逸
}

func receiveInterface(i interface{}) {
	fmt.Println(i)
	_ = i
}

func addArray(a [100000]int) {
	for i, v := range a {
		a[i] = v + 1
	}
}

func printArray2(b *[100000]int) {
	//fmt.Println(a)
}
