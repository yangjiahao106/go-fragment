package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

const (
	CommandConn   = iota + 0x01 // 0x01，连接请求包
	CommandSubmit               // 0x02，消息请求包
)

const (
	CommandConnAck   = iota + 0x81 // 0x81，连接请求的响应包
	CommandSubmitAck               // 0x82，消息请求的响应包
)

// tcp-server-demo1/frame/frame.go

var ErrShortWrite = errors.New("short write")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct{}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

type FramePayload []byte
type StreamFrameCodec interface {
	Encode(io.Writer, FramePayload) error   // data -> frame，并写入io.Writer
	Decode(io.Reader) (FramePayload, error) // 从io.Reader中提取frame payload，并返回给上层
}

func (p *myFrameCodec) Encode(w io.Writer, framePayload FramePayload) error {
	var f = framePayload
	var totalLen int32 = int32(len(framePayload)) + 4

	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}

	n, err := w.Write([]byte(f)) // write the frame payload to outbound stream
	if err != nil {
		return err
	}

	if n != len(framePayload) {
		return ErrShortWrite
	}

	return nil
}

func (p *myFrameCodec) Decode(r io.Reader) (FramePayload, error) {
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, totalLen-4)

	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}

	return FramePayload(buf), nil
}

// tcp-server-demo1/packet/packet.go

type Packet interface {
	Decode([]byte) error     // []byte -> struct
	Encode() ([]byte, error) //  struct -> []byte
}

// tcp-server-demo1/packet/packet.go

type Submit struct {
	ID      string
	Payload []byte
}

func (s *Submit) Decode(pktBody []byte) error {
	s.ID = string(pktBody[:8])
	s.Payload = pktBody[8:]
	return nil
}

func (s *Submit) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), s.Payload}, nil), nil
}

type SubmitAck struct {
	ID     string
	Result uint8
}

func (s *SubmitAck) Decode(pktBody []byte) error {
	s.ID = string(pktBody[0:8])
	s.Result = uint8(pktBody[8])
	return nil
}

func (s *SubmitAck) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), []byte{s.Result}}, nil), nil
}

// tcp-server-demo1/packet/packet.go

func PacketDecode(packet []byte) (Packet, error) {
	commandID := packet[0]
	pktBody := packet[1:]

	switch commandID {
	case CommandConn:
		return nil, nil
	case CommandConnAck:
		return nil, nil
	case CommandSubmit:
		s := SubmitPool.Get().(*Submit)
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return s, nil
	case CommandSubmitAck:
		s := SubmitAck{}
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return &s, nil
	default:
		return nil, fmt.Errorf("unknown commandID [%d]", commandID)
	}
}

// tcp-server-demo1/packet/packet.go

func PacketEncode(p Packet) ([]byte, error) {
	var commandID uint8
	var pktBody []byte
	var err error

	switch t := p.(type) {
	case *Submit:
		commandID = CommandSubmit
		pktBody, err = p.Encode()
		if err != nil {
			return nil, err
		}
	case *SubmitAck:
		commandID = CommandSubmitAck
		pktBody, err = p.Encode()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown type [%s]", t)
	}
	return bytes.Join([][]byte{[]byte{commandID}, pktBody}, nil), nil
}

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p Packet
	p, err = PacketDecode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error:", err)
		return
	}

	switch p.(type) {
	case *Submit:
		submit := p.(*Submit)
		fmt.Printf("recv submit: id = %s, payload=%s\n", submit.ID, string(submit.Payload))
		submitAck := &SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		SubmitPool.Put(submit)
		ackFramePayload, err = PacketEncode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := NewMyFrameCodec()

	// 增加缓冲
	readerWithBuf := bufio.NewReader(c)
	writerWithBuf := bufio.NewWriter(c)

	for {
		// decode the frame to get the payload
		framePayload, err := frameCodec.Decode(readerWithBuf)
		if err != nil {
			fmt.Println("handleConn: frame decode error:", err)
			return
		}

		// do something with the packet
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error:", err)
			return
		}

		// write ack frame to the connection
		err = frameCodec.Encode(writerWithBuf, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
	}
}


var SubmitPool = sync.Pool{
	New: func() interface{} {
		return &Submit{}
	},
}

func main() {

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	//// 绑定一个handler
	//http.Handle("/", http.StripPrefix("/static/", http.FileServer(http.Dir("./output"))))
	//// 监听服务
	//_ = http.ListenAndServe(":8000", nil)
	go func() {
		err = http.ListenAndServe("127.0.0.1:6060", nil)
		if err != nil {
			log.Fatal("funcRetErr=http.ListenAndServe")
		}
	}()


	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				fmt.Println("accept error:", err)
				break
			}
			// start a new goroutine to handle the new connection.
			go handleConn(c)
		}
	}()

	startClient(1)

}
