package channel_sample

import (
	"fmt"
	"time"
)

// Counter 代表计数器的类型。
type Counter struct {
	count int
}

var mapChan2 = make(chan map[string]Counter, 1)

func Chan_With_Map2() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan2; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		countMap := map[string]Counter{
			"count": {},
		}
		for i := 0; i < 5; i++ {
			mapChan2 <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan2)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
