package main

import "runtime"

func GoPark() {
	go func() {
		runtime.Gosched()
		runtime.Goexit()

	}()

}
