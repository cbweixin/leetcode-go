package channel_sample

import (
	"fmt"
	"time"
)

func Timer_test2() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	select {
	case e := <-intChan:
		fmt.Printf("Received: %v\n", e)
	case <-time.NewTimer(time.Millisecond * 500).C:
		fmt.Println("Timeout!")
	}
}
