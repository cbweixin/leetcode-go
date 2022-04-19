package channel_sample

import (
	"fmt"
	"time"
)

var mapChan3 = make(chan map[string]*Counter, 1)

func Chan_With_Map3() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan3; ok {
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
		// Counter is a pointer
		// compare with chanel_ex4.go, now Counter is a pointer, then its status can be mutated
		countMap := map[string]*Counter{
			// get the address of counter, ie reference
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan3 <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan3)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func (counter *Counter) String() string {
	return fmt.Sprintf("{count : %d}", counter.count)
}

/**
The count map: map[count:{count : 1}]. [sender]
The count map: map[count:{count : 2}]. [sender]
The count map: map[count:{count : 3}]. [sender]
The count map: map[count:{count : 4}]. [sender]
The count map: map[count:{count : 5}]. [sender]
Stopped. [receiver]
*/
