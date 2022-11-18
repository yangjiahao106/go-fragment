package main

import (
	"fmt"
	"runtime"
	"time"
)

func Call() {
	var i = 1
	_ = i
}

func TestOneProcessor() {

	runtime.GOMAXPROCS(1)

	go func() {
		for {
			//time.Sleep(time.Millisecond)
			i := 0
			for ; i < 100000; i++ {
				Call()
			}
			fmt.Println("A", time.Now().UnixMilli())
		}
	}()

	go func() {
		for {
			//time.Sleep(time.Millisecond)
			i := 0
			for ; i < 100000; i++ {
				Call()
			}
			fmt.Println("B", time.Now().UnixMilli())
		}
	}()

	//go func() {
	//	for {
	//		for i := 0; i < 100000; i++ {
	//		}
	//		//time.Sleep(time.Millisecond)
	//		fmt.Println(3)
	//	}
	//}()

	time.Sleep(time.Millisecond * 30)

}
