package main

import (
	"fmt"
	"net"
	"time"
)

func startClient(i int) {
	quit := make(chan struct{})
	done := make(chan struct{})
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	// shutdown
	//defer conn.(*net.TCPConn).CloseWrite()
	defer conn.Close()

	fmt.Printf("[client %d]: dial ok", i)

	frameCodec := NewMyFrameCodec()
	var counter int


	//  handle ack
	go ackHandler(i, quit, done, conn, frameCodec)

	for {
		// send submit
		counter++
		id := fmt.Sprintf("%08d", counter) // 8 byte string

		s := &Submit{
			ID:      id,
			Payload: []byte("payload"),
		}

		framePayload, err := PacketEncode(s)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[client %d]: send submit id = %s, payload=%s, frame length = %d\n",
			i, s.ID, s.Payload, len(framePayload)+4)

		err = frameCodec.Encode(conn, framePayload)
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
		if counter >= 100 {
			quit <- struct{}{}
			<-done
			fmt.Printf("[client %d]: exit ok", i)
			return
		}
	}
}

func ackHandler(i int, quit chan struct{}, done chan struct{}, conn net.Conn, frameCodec StreamFrameCodec) {
	func() {
		// handle ack
		for {
			select {
			case <-quit:
				done <- struct{}{}
				return
			default:
			}

			_ = conn.SetReadDeadline(time.Now().Add(time.Second * 1))

			ackFramePayLoad, err := frameCodec.Decode(conn)
			if err != nil {
				if e, ok := err.(net.Error); ok {
					if e.Timeout() {
						continue
					}
				}
				panic(err)
			}

			p, err := PacketDecode(ackFramePayLoad)
			submitAck, ok := p.(*SubmitAck)
			if !ok {
				panic("not submitack")
			}
			fmt.Printf("[client %d]: the result of submit ack[%s] is %d\n", i, submitAck.ID, submitAck.Result)
		}
	}()
}
