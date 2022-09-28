package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGPIPE)
		sig := <-ch
		fmt.Println("receive signal:", sig)
	}()

	c, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}

	time.Sleep(time.Second * 5)

	for i := 0; i < 10000; i++ {
		fmt.Println("start write")
		n, err := c.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			fmt.Println("Write error:", err, n)
			return
		}
		time.Sleep(time.Second)
		buf := make([]byte, 100)
		n, err = c.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err, n)
			return
		}

		fmt.Println("write success: ", i, n)
		time.Sleep(time.Second * time.Duration(i))
	}

	time.Sleep(time.Second)
}
