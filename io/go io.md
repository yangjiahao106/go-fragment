go io

1. 基本的 IO 接口

io包中的核心接口io.Reader io.Writer

//Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。
//即使 Read 返回的 n < len(p)，它也会在调用过程中占用 len(p) 个字节作为暂存空间。若可读取的数据
//不到 len(p) 个字节，Read 会返回可用数据，而不是等待更多数据。
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Write 将 len(p) 个字节从 p 中写入到基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）
// 以及任何遇到的引起写入提前停止的错误。若 Write 返回的 n < len(p)，它就必须返回一个 非nil 的错误。
type Writer interface {
	Write(p []byte) (n int, err error)
}

2. 实现了Reader Write接口的类型

bytes.Buffer是集读写功能于一身的数据类型,它非常适合作为字节序列的缓冲区。它的指针
类型实现的接囗就更多了。更具体地说,该指针类型实现的读取相关的接囗有下面几个。

1. io.Reader ;
2. io.ByteReader ;
3. io.RuneReader ;
4. io.ByteScanner;
5. io.RuneScanner;
6. io.WriterTo ;

其实现的写入相关的接囗则有

1. io.Writer
2. io.ByteWriter
3. io.ReaderFrom 

那么,这些类型实现了这么多的接囗,其动机或者说目的究竟是什么呢?

简单地说，这是为了提高不同程序实体之间的互操作性。
在io包中，有这样几个用于拷贝数据的函数：

func Copy(dst Writer, src Reader) (written int64, err error) 
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) 
func CopyN(dst Writer, src Reader, n int64) (written int64, err error) 

虽然这几个函数在功能上都略有差别，但是它们都首先会接受两个参数，即：用于代表数据目的地、io.Writer类型的参数dst，
以及用于代表数据来源的、io.Reader类型的参数src。这些函数的功能大致上都是把数据从src拷贝到dst。
不论我们给予它们的第一个参数值是什么类型的，只要这个类型实现了io.Writer接口即可。
同样的，无论我们传给它们的第二个参数值的实际类型是什么，只要该类型实现了io.Reader接口就行。
一旦我们满足了这两个条件，这些函数几乎就可以正常地执行了。

func TestCopyN(t *testing.T) {
	src := strings.NewReader("CopyN copies n bytes (or until an error) from src to dst. " +
		"It returns the number of bytes copied and the earliest error encountered while copying.")

	dst := new(strings.Builder)

	written, err := io.CopyN(dst, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("Written to dst (%d): %q\n", written, dst.String())

	dst2 := bytes.NewBuffer(nil)
	written, err = io.CopyN(dst2, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("Written to dst2(%d): %q\n", written, dst.String())

}


变量src和dst的类型分别是strings.Reader和strings.Builder，
但是当它们被传到io.CopyN函数的时候，就已经分别被包装成了io.Reader类型和io.Writer类型的值。
io.CopyN函数也根本不会去在意，它们的实际类型到底是什么。
如此一来，Go 语言的各种库中，能够操作它们的函数和数据类型明显多了很多。


2. bufio

顾名思义， bufio包中实现的IO操作都内置了缓冲区
主要类型：
Reader
Writer 
Scanner 
ReadWriter 

缓冲读

// 默认缓冲区大小
const (
    defaultBufSize = 4096
)

// 最小缓冲区大小 自定义小于此阈值将会被覆盖
const minReadBufferSize = 16

// 使用默认缓冲区大小
bufio.NewReader(rd io.Reader)
// 使用自定义缓冲区大小
bufio.NewReaderSize(rd io.Reader, size int)




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
	reader := bufio.NewReaderSize(file, 200) // defaultBufSize = 4096
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		line := make([]byte, 81)
		_, err := reader.Read(line) //读取数量可能小于len(line)
		//_, err := io.ReadFull(reader, line)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(line))
	}
	file.Close()
}


cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkRead-16                  100000              1209 ns/op
BenchmarkReadWithBufIO-16         100000                84.89 ns/op
PASS



Read(p []byte) (n int, err error)

Read方法缓冲读的大致过程如下，设定好缓冲区大小buf_size后，读取的字节数为len(p)，缓冲的字节数为bn：

1. 如果缓冲区为空，且 len(p) >= buf_size，则直接从文件读取，不启用缓冲。
2. 如果缓冲区为空，且 len(p) < buf_size，则从文件读取buf_size 字节的内容到缓冲区，程序再从缓冲区中读取len(p)字节的内容，此时缓冲区剩余bn = buf_size - len(p)字节。
3. 如果缓冲区不为空，len(p) < bn，则从缓冲区读取len(p)字节的内容，不发生文件IO。
4. 如果缓冲区不为空，len(P) >= bn，则从缓冲区读取bn字节的内容，不发生文件IO，缓冲区置为空，此时读取的字节数量n小于len(p)

如果要精确读取len(p)字节，请使用io.ReadFull

缓冲读通过预读，可以在一定程度上减少文件IO次数，故提高性能。



缓冲写:

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

	w.Flush() // 刷新缓冲区剩余的数据
	file.Close()
}

cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkWrite-16                 100000             10773 ns/op
BenchmarkWriteWithBufIO-16        100000               289.3 ns/op
PASS


缓冲写的大致过程如下，设定好缓冲区大小buf_size后，写入的字节数为wn，缓冲的字节数为bn：

如果 wn + bn < buf_size，则程序将内容写入缓冲区，不发生文件IO。
如果 wn + bn > buf_size 且缓冲区为空，则直接写入文件，不启用缓冲，发生文件IO
如果 wn + bn > buf_size 且缓冲区不为空，则先将缓冲取写满，并刷新缓冲区，之后跳到第二步

简单说就是要写入的内容先缓冲着，缓冲不下了则将缓冲区内容写入文件。

