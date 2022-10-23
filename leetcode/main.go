package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	now = time.Now().Add(time.Hour * 24 * 14) // 调整活动时间， 上线删除

	fmt.Println(now)

	TestRemain()
	//testFindKthLargest()
}
