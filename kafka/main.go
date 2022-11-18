package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//producer(0)
	//producer(1)

	//consumer()
	c := make(chan os.Signal, 1)
	go signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	Reader(1, c)
	Writer(c)
	time.Sleep(time.Second )
	//go Reader(0)
	//time.Sleep(time.Second * 5)
	//go producer(0)
	//time.Sleep(time.Second * 10)
	//go Reader(2)

	//time.Sleep(time.Hour)
	//consumer()

}


func fileReadAt() {
	file, _ := os.Open("")
	//file.ReadAt()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

}

var topic = "test_topic"

func Reader(id int, c chan os.Signal) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"apitest.aishepin8.com:9092", "apitest.aishepin8.com:9093", "apitest.aishepin8.com:9094"},
		GroupID: "consumer-group-id",
		//Partition:   0,
		Topic:       topic,
		MaxWait:     time.Second,
		MinBytes:    1000, //
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.FirstOffset,
		//GroupBalancers: []kafka.GroupBalancer{kafka.RoundRobinGroupBalancer{}},
		//CommitInterval: // 控制异步提交
		//ReadLagInterval: -1,
	})

	var err error
	//err = r.SetOffset(337)
	//err := r.SetOffsetAt(context.Background(), time.Now().Add(-time.Hour*10))
	//fmt.Println(err)

	//r.SetOffset(313)

	lag, err := r.ReadLag(context.Background())
	if err != nil {
		fmt.Println("lag err", err)
	}
	fmt.Println("lag: ", lag)


	for i := 0; ; i++ {
		if len(c) > 0 {
			fmt.Println("break")
			break
		}
		//m, err := r.ReadMessage(context.Background()) // 自动提交偏移量
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			fmt.Println("fetchMessage err:", err)
			break
		}

		fmt.Printf("ID%v: message at topic/partition/offset %v/%v/%v: %s = %s\n", id, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		err = r.CommitMessages(context.Background(), m) //提交偏移量
		if err != nil {
			fmt.Println("commit message failed ", err)
			break
		}
	}
	//time.Sleep(time.Second * 4)
	fmt.Println("start close")
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func Writer(c chan os.Signal) {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:         kafka.TCP("apitest.aishepin8.com:9092", "apitest.aishepin8.com:9093", "apitest.aishepin8.com:9094"),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		BatchSize:    5,
		//BatchBytes: // default 1MB
		BatchTimeout: time.Second / 100, // 默认1s
		//Completion:    // 回调函数
		MaxAttempts: 10, // 默认重试10次
		Async:       true,
		Completion: func(messages []kafka.Message, err error) {
			if err == nil {
				fmt.Println("batch Send success, BachSize:", len(messages))
			} else {
				fmt.Println("batch Send failed, BachSize:", len(messages), err)
			}
		},
	}


	i := 1
	for {
		i++
		if len(c) > 0 {
			fmt.Println("break")
			break
		}

		messages := []kafka.Message{
			{
				Key:   []byte(fmt.Sprintf("Key:%v", i)),
				Value: []byte(fmt.Sprintf("Val:%v:", i)),
			},
		}
		fmt.Println("a")
		err := w.WriteMessages(context.Background(), messages...)
		if err != nil {
			writeErr := err.(kafka.WriteErrors)
			for i := 0; i < len(messages); i++ {
				if writeErr[i] != nil {
					fmt.Println("sendMessageFailed, message:", messages[i])
				}
			}
		}

		fmt.Println("write success ", i)
		time.Sleep(time.Second / 10)
		if err != nil {
			fmt.Println("failed to write messages:", err)
			break
		}

	}
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	fmt.Println("closed!")

}

func producer(partition int) {

	//partition := 1

	conn, err := kafka.DialLeader(context.Background(), "tcp", "apitest.aishepin8.com:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal("SetWriteDeadline", err)
	}

	for i := 0; i < 10; i++ {
		i, err := conn.WriteMessages(
			kafka.Message{Key: []byte("Key"), Value: []byte(fmt.Sprintf("Msg:%v Partition %v", i, partition))},
			//kafka.Message{Key: []byte("two"), Value: []byte("two!")},
			//kafka.Message{Key: []byte("three"), Value: []byte("three!")},
		)
		if err != nil {
			if err == kafka.LeaderNotAvailable || err == kafka.NotLeaderForPartition {
				continue
			}
			log.Fatal("failed to write messages:", err)
		}
		fmt.Println("send success", i)

	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	fmt.Printf("finished")
}

func consumer() {
	// to consume messages
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "apitest.aishepin8.com:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	fmt.Println(conn.Offset())

	offset, err := conn.Seek(240, 1)
	fmt.Println(offset, err)
	fmt.Println(conn.Offset())
	_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	batch := conn.ReadBatch(10e2, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {

		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))

	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
