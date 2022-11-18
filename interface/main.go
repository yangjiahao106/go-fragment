package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"
)

type MyInterface interface {
	M1()
}

type T int

func (T) M1() {
	println("T's M1")
}

func (T) M2() {
	println("T's M2")
}
func (T) Read(p []byte) (n int, err error) {
	return 0, nil
}

type MyReader struct {
	io.Reader       // underlying reader
	N         int64 // max bytes remaining
	t         T
}

//func (r *MyReader) Read(p []byte) (n int, err error) {
//	fmt.Println("my read")
//	return 0, nil
//}

func server(f func(w http.ResponseWriter, r *http.Request)) {
	f(nil, nil)
}

const (
	a = iota
	b
	c
	d
)

func main() {


	fmt.Println(a, b, c, d)

	s := "你好吗kk"
	fmt.Println(len(s), utf8.RuneCountInString(s))
	fmt.Println(string([]rune(s)[1:]))
	for k, v := range s {
		fmt.Println(k, string(v))
	}

	//bytes.NewBuffer()
	TestCapReader()
	r := MyReader{
		Reader: bytes.NewBuffer(nil),
	}
	_, _ = r.Read(nil)

	bytes.NewBuffer(nil)

	return
	var t T
	var i interface{} = t
	v1, ok := i.(MyInterface)
	if !ok {
		panic("the value of i is not MyInterface")
	}
	v1.M1()
	fmt.Printf("V:%t\n", t)
	fmt.Printf("the type of v1 is %T\n", v1) // the type of v1 is main.T 打印的是动态类型

	i = int64(13)
	v2, ok := i.(MyInterface)
	fmt.Printf("the type of v2 is %T\n", v2) // the type of v2 is <nil>
	// v2 = 13 //  cannot use 1 (type int) as type MyInterface in assignment: int does not implement MyInterface (missing M1   method)
}

type Error struct {
}

func (e *Error) Error() string {
	return ""
}

// 关于 nil != nil
func foo() {
	var err error
	var err2 error

	err2 = (*Error)(nil)
	fmt.Println(err == err2)

	if err2 != nil {
		fmt.Println("err != nil ")
	}
}

// 接口复制也是值拷贝
func interfaceCopy() {
	var i interface{}
	println(i)
	fmt.Println(i == nil) // true
	i = (*Error)(nil)
	println(i)
	fmt.Println(i == nil) // false

	s := [3]int{1, 2, 3}
	i = s // 值拷贝
	fmt.Println(i)
	s[0] = 0
	fmt.Println(i)

}

// 装饰器模式
func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

type capitalizedReader struct {
	r io.Reader
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}
	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

func TestCapReader() {
	r := strings.NewReader("hello, gopher!\n")
	//io.LimitReader(r, 4)
	r1 := CapReader(r)
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}
