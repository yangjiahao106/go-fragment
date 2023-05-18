package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestContextWithCancel(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(dst)
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 可以 重复cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel() // 这里为了使不熟悉 go 的更能明白在这里调用了 cancel()
			break
		}
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

// parent 取消了 children 也会取消
func TestContextWithCancel2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 可以 重复cancel

	ctx2, cancel2 := context.WithCancel(ctx)
	defer cancel2() // 可以 重复cancel
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_ = <-ctx2.Done()
		fmt.Println("child cancel")
		wg.Done()
	}()
	fmt.Println("parent cancel")
	cancel()
	wg.Wait()
}

func TestContextWithTimeout(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(dst)
					return // returning not to leak the goroutine
				case dst <- n:
					fmt.Println(time.Now().UnixMilli())

					time.Sleep(time.Second / 10)
					n++
				}
			}
		}()
		return dst
	}

	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	fmt.Println(time.Now().UnixMilli())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
	}

}

func TestContextWithValue(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(dst)
					return // returning not to leak the goroutine
				case dst <- n:
					time.Sleep(time.Second / 10)
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
	}

}
