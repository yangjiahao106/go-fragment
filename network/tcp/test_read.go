package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				fmt.Println("accept error:", err)
				break
			}
			// start a new goroutine to handle the new connection.
			go func(conn net.Conn) {
				buffer := make([]byte, 1)
				//n, err := io.ReadFull(c, buffer)
				for {
					time.Sleep(time.Second * 3)
					n, err := conn.Read(buffer)
					if err != nil {
						fmt.Println("ReadFull error:", err)
						return
					}
					fmt.Println("read success: ", n, string(buffer))
				}

			}(c)
		}
	}()

	c, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}

	for i := 0; i < 10; i++ {
		n, err := c.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
		fmt.Println("write success: ", i, n)
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second)
}
