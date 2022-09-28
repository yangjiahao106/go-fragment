package main

import (
	"fmt"
	"reflect"
	"time"
)

func main2() {
	var ch = make(chan int, 10)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(time.Second * 3)
		close(ch)
	}()

	go func() {
		fmt.Println("go exit")
		close(ch)
	}()
	ch = nil
	fmt.Println("finished")
	time.Sleep(time.Second * 2)

	//for i := 0; i < 10; i++ {
	//	select {
	//	case ch <- i:
	//	case v := <-ch:
	//		fmt.Println(v)
	//	}
	//}
	//time.Sleep(time.Second)
}

// 关闭 channel
func CloseChan() {
	c := make(chan int, 10)
	c <- 1
	close(c)
	fmt.Println(<-c) // 返回未读取的值
	fmt.Println(<-c) // 返回零值

	c <- 1 // panic

}

// 测试 扇入模式
func TestFinInReflect() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	out := fanInReflect(c1, c2, c3)
	go func() {
		c1 <- 1
		c3 <- 3
		c2 <- 2
		close(c1)
		close(c2)
		close(c3)

	}()

	for {
		select {
		case v, ok := <-out:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}

}

// 扇入模式 reflect
func fanInReflect(ins ...chan interface{}) chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		cases := make([]reflect.SelectCase, 0)
		for _, c := range ins {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				// closed
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}

			out <- v.Interface()
		}
	}()

	return out
}

func testReflectSelect() {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建SelectCase
	var cases = createCases(ch1, ch2)

	// 执行10次select
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}

	return cases
}

func TestReduce() {
	ch := make(chan interface{}, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
		close(ch)
	}()

	out := reduce(ch, func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	})

	fmt.Println(out)

}

func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil { // 异常检查
		return nil
	}

	out := <-in         // 先读取第一个元素
	for v := range in { // 实现reduce的主要逻辑
		out = fn(out, v)
	}

	return out
}

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

// 测试 内存重排
func rearrangement() {
	go setup()
	for !done {
	}
	print(a)
}



type T struct {
	msg string
}

var g *T

func setup2() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func rearrangement2() {
	go setup2()
	for g == nil {
    }
	print(g.msg)
}