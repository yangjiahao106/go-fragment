package main

func TestSelect() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})

	for {
		select {
		case v1, ok := <-c1:
			_ = v1
			// 如果c1被关闭(ok==false)，每次从c1读取都会立即返回，将导致死循环
			// 可以通过将c1置为nil来让select ignore掉这个case，继续评估其它case
			if !ok {
				c1 = nil
			}

		case v2 := <-c2:
			_ = v2
		// 同样，如果c2被关闭，每次从c1读取都会立即返回对应元素类型的零值(如空字符串)，导致死循环
		// 解决方案仍然是置c2为nil，但是有可能误判(写入方是写入了一个零值而不是关闭channel，比如整数0)

		case c3 <- 1:
			// 如果c3已经关闭，则panic
			// 如果c3为nil，则ignore该case
		}

	}

}
