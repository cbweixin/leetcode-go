package channel_sample

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func Chan_With_Map() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	// what would happen if I remove one <-synChan? below is what I got
	// The count map: map[count:1]. [sender]
	// The count map: map[count:2]. [sender]
	// The count map: map[count:3]. [sender]
	// The count map: map[count:4]. [sender]
	// The count map: map[count:5]. [sender]

	// where is "Stopped. [receiver]"?
	// we have : 	syncChan := make(chan struct{}, 2)
	// both sender and reciever would do this : syncChan <- struct{}{}
	// so when the main program would call <-synChan twice, the main program would blocked
	// and wait for both reciever and sender to finish its task, if we have only one <-synChan
	// then main program would not wait both, most likey the send do syncChan <- struct{}{}, then the main program
	// just exit, so we will never saw "Stopped. [receiver] "
	<-syncChan
	<-syncChan

}

/**
The count map: map[count:1]. [sender]
The count map: map[count:2]. [sender]
The count map: map[count:3]. [sender]
The count map: map[count:4]. [sender]
The count map: map[count:5]. [sender]
Stopped. [receiver]
*/
