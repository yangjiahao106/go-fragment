package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var SERVER_RECV_LEN = 1000

func main() {
	go server()

	client()
}


func server() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("ResolveUDPAddr error:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	defer conn.Close()

	for {
		data := make([]byte, 1024)
		n, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("Read From Udp err =", err)
			continue
		}
		strData := string(data)
		fmt.Println("Server Received:", strData, n)

		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			fmt.Println("Write Udp err =", err)
			continue
		}
		fmt.Println("Server Response:", upper)
	}

}


func client() {
	conn, err := net.Dial("udp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > SERVER_RECV_LEN {
				toWrite = line[written : written+SERVER_RECV_LEN]
			} else {
				toWrite = line[written:]
			}

			n, err = conn.Write([]byte(toWrite))
			if err != nil {
				fmt.Println("client write error:", err)
				return
			}
			fmt.Println("Client Write:", toWrite)

			msg := make([]byte, SERVER_RECV_LEN)
			n, err = conn.Read(msg)
			if err != nil {
				fmt.Println("client read error:", err)
				return
			}

			fmt.Println("Client Revc:", string(msg))
		}
	}

}

