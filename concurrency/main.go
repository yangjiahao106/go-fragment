package main

import (
	"fmt"
	"time"
)

func main()  {
	//rearrangement2()

	t := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-t.C:
			fmt.Println("f")
		}
	}

}