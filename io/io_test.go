package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestIo(t *testing.T) {
	//io.CopyBuffer()
	//io.CopyN()
	//os.Stdout
	//io.ReadWriter()
	//os.ReadFile()
	fmt.Println()
	//io.ReadAll()
	//bytes.Buffer{}
	//bufio.NewReader()
	//io.Reader()
	//io.Writer()
	//io.Closer()
	//io.ReaderFrom()

}

func TestCopyN(t *testing.T) {
	src := strings.NewReader("CopyN copies n bytes (or until an error) from src to dst. " +
		"It returns the number of bytes copied and the earliest error encountered while copying.")

	dst1 := new(strings.Builder)

	written, err := io.CopyN(dst1, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("Written to dst1(%d): %q\n", written, dst1.String())

	dst2 := bytes.NewBuffer(nil)
	written, err = io.Copy(dst2, src)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("Written to dst2(%d): %q\n", written, dst2.String())

}

func FileToByteArray(fileName string) []byte {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)

	//bufio.NewScanner()
	if _, err = io.Copy(buf, file); err != nil {
		return nil
	}

	return buf.Bytes()
}

func TestWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	TimeGenerator(buf)
	fmt.Println(buf.String())

	TimeGenerator(os.Stdout)

}

func TimeGenerator(writer io.Writer) {

	_, _ = writer.Write([]byte(time.Now().Format("2006-01-02 15:04:05.000000000")))

	_, _ = writer.Write([]byte{10})

}

// 写入 go test -bench='BenchmarkWrite|BenchmarkWriteWithBufIO' -run=none -benchtime=100000x
func BenchmarkWrite(b *testing.B) {
	file, err := os.Create(fmt.Sprintf("./tmp_%d.txt", time.Now().UnixMilli()))
	if err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := file.WriteString("Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	file.Close()
}

func BenchmarkWriteWithBufIO(b *testing.B) {
	file, err := os.Create(fmt.Sprintf("./tmp_%d.txt", time.Now().UnixMilli()))
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	w := bufio.NewWriter(file)
	for i := 0; i < b.N; i++ {
		_, err := w.WriteString("Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()
	file.Close()
}

// 读取
// go test -bench='BenchmarkReadWithBufIO|BenchmarkRead' -run=none -benchtime=100000x
func BenchmarkRead(b *testing.B) {
	file, err := os.Open("./tmp.txt")
	if err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		line := make([]byte, 81)
		_, err := file.Read(line)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(line))
	}
	file.Close()
}

func BenchmarkReadWithBufIO(b *testing.B) {
	file, err := os.Open("./tmp.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file) // defaultBufSize = 4096
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		line := make([]byte, 81)
		//_, err := reader.Read(line) //读取数量可能小于len(line)
		_, err := io.ReadFull(reader, line)
		//reader.ReadBytes('\n')
		//reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(line))
	}
	file.Close()
}

func BenchmarkScanner(b *testing.B) {
	file, err := os.Open("./tmp.txt")
	if err != nil {
		log.Fatal(err)
	}
	scnner := bufio.NewScanner(bufio.NewReader(file)) // defaultBufSize = 4096
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if scnner.Scan() {
			fmt.Println(string(scnner.Bytes()))
		}
	}
	file.Close()

}
