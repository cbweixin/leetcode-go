package channel_sample

import (
	"fmt"
	"time"
)

var strChan4 = make(chan string, 3)

func Single_Direction_Channel() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan4, syncChan1, syncChan2) // 用于演示接收操作。
	go send(strChan4, syncChan1, syncChan2)    // 用于演示发送操作。
	<-syncChan2
	<-syncChan2
}

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
	close(strChan4)
	syncChan2 <- struct{}{}
}
