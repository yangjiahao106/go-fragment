package main

import (
	"encoding/binary"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func h() []*int {

	s := []*int{new(int), new(int), new(int), new(int)}
	// do something with s ...

	// Reset pointer values.
	s[0], s[len(s)-1] = nil, nil
	fmt.Println(s[:])
	fmt.Println(s[1:3], cap(s[1:3:4]))
	return s[1:3:4]
}

func TestSyncMap (){
	var i interface{} = nil
	p  := unsafe.Pointer(&i)

	fmt.Println(p)
	sm := sync.Map{}
	sm.Store("a", nil)
	k, ok := sm.Load("a")
	fmt.Println(k, ok)
	return

}

func main() {

	TestSyncMap()
	return

	//fmt.Println( h())
	// go 内存泄露， 垃圾回收触发条件, 内存逃逸的情况
	// 排序算法

	//runtime.GOMAXPROCS() P 的数量
	//debug.SetMaxThreads() M 的最大数量
	ticker := time.NewTicker(time.Second)
	i := 0
	for true {
		select {
		case <-ticker.C:
			fmt.Println(i)
			i++
		}
	}

	fmt.Println(binary.BigEndian.Uint16([]byte{1, 0}))
	fmt.Println(binary.LittleEndian.Uint16([]byte{1, 0}))

	m := map[string]interface{}{}

	m["1"] = 1
	if m["1"] == 1 {
		fmt.Println("1***")
	}

	fmt.Println("hello world")
	{
		foo := NewFoo(WithName("name"), WithAge(1), WithDB("db"))
		fmt.Println(foo)

	}
	foo := "xx"
	fmt.Println(foo)
}

type Foo struct {
	name string
	age  int
	db   interface{}
}

type Option func(foo *Foo)

func WithName(name string) Option {
	return func(foo *Foo) {
		foo.name = name
	}
}

func WithAge(age int) Option {
	return func(foo *Foo) {
		foo.age = age
	}
}

func WithDB(db interface{}) Option {
	return func(foo *Foo) {
		foo.db = db
	}
}

func NewFoo(options ...Option) *Foo {
	foo := &Foo{}
	for _, op := range options {
		op(foo)
	}
	return foo
}
