package main

import (
	"fmt"

	"unsafe"
)



func main() {
	fmt.Println(unsafe.Sizeof(1))
	var arr = [5]int{11, 12, 13, 14, 15}
	var p *int = &arr[0]
	var i uintptr
	for i = 0; i < uintptr(len(arr)); i++ {
		p1 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i*uintptr(8)))
		println(*p1)
	}

	bs := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b := (*byte)(unsafe.Pointer(&bs))
	fmt.Println(*b)
	by := byte('2')

	fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&bs)) + unsafe.Sizeof(by))))
}
