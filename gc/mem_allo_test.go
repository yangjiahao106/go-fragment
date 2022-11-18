package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMem(t *testing.T) {

	type demo1 struct {
		a int8
		b int16
		c int32
	}

	type demo2 struct {
		a int8
		c int32
		b int16
	}

	type demo3 struct {
		a int8
		b int16
		c int64
	}

	type demo4 struct {
		a int8
		c int64
		b int16
	}

	type demo5 struct {
		a int8
		c int64
		//b int16
		d struct{} // 空结构体作为最后一个字段需要内存对齐
	}


	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12
	fmt.Println(unsafe.Sizeof(demo3{})) // 16
	fmt.Println(unsafe.Sizeof(demo4{})) // 24
	fmt.Println(unsafe.Sizeof(demo5{}))


}
