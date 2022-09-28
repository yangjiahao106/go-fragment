package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("started")
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

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
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println("ReadFull error:", err)
					return
				}
				fmt.Println("read success: ", n, string(buffer))
				conn.(*net.TCPConn).CloseWrite()
			}

		}(c)
	}
}
