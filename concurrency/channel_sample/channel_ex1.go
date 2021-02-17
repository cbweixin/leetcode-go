package channel_sample

import (
	"fmt"
	"time"
)

func Channel_test() {
	exit := make(chan bool)

	go func() {
		for {
			select {
			case <-exit:
				fmt.Println("退出监控")
				return
			default:
				fmt.Println("监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("通知监控退出")
	exit <- true

	//防止main goroutine过早退出
	time.Sleep(5 * time.Second)
}
