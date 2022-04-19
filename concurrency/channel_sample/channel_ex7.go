package channel_sample

import (
	"fmt"
	"time"
)

// bi-directional channel
var strChan4 = make(chan string, 3)

func Single_Direction_Channel() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan4, syncChan1, syncChan2) // 用于演示接收操作。
	go send(strChan4, syncChan1, syncChan2)    // 用于演示发送操作。
	<-syncChan2
	<-syncChan2
}

// notice, single direction channel, be aware of chan and <- position
// <- chan means recieve, chan <- means send, `strchan4` is bi-directional channel,
// but in this method, parameter strChan4 would be required as a singel directional chan
func receive(strChan4 <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan4; ok {
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}

	// you can not close a recieving channel, compile error
	// close(strChan4)
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan4 chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan4 <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second * 2)
	// strChan4 is sending channel, so it could be closed. but you can not close recieve channel, this would lead to
	// a compile error
	close(strChan4)
	syncChan2 <- struct{}{}
}

/**
Sent: a [sender]
Sent: b [sender]
Sent: c [sender]
Sent a sync signal. [sender]
Received a sync signal and wait a second... [receiver]
Received: a [receiver]
Received: b [receiver]
Received: c [receiver]
Received: d [receiver]
Sent: d [sender]
Wait 2 seconds... [sender]
Stopped. [receiver]

Program exited.
*/
