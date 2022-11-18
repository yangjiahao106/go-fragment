package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var once3 = sync.Once{}
var instance *singletonInstance

type singletonInstance struct {
}

// once 实现单例模式
func newInstance() *singletonInstance {
	once3.Do(
		func() {
			instance = &singletonInstance{}
		})
	return instance
}

var initialized int32
var mu = sync.Mutex{}

// mutex 实现单例模式
func newInstance2() *singletonInstance {
	if atomic.LoadInt32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()
	if atomic.LoadInt32(&initialized) == 0 {
		instance = &singletonInstance{}
		atomic.StoreInt32(&initialized, 1)
	}
	return instance
}

func TestOnce(t *testing.T) {

	once := sync.Once{}

	once.Do(func() {
		fmt.Println("once")
	})

	once.Do(func() {
		fmt.Println("once2")
	})

	once2 := sync.Once{}
	once2.Do(func() {
		fmt.Println("once2")
	})

}
