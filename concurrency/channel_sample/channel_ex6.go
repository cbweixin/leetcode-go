package channel_sample

import "fmt"

func Close_test() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		<-syncChan1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		// firstly close the channel
		// channel can only be closed once, repeatly closing would cause panic
		// if channel variable become nil, it would also caused panic.
		close(dataChan)
		// then start the receiver, but receiver still able to get all element in channel
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

/**
Sent: 0 [sender]
Sent: 1 [sender]
Sent: 2 [sender]
Sent: 3 [sender]
Sent: 4 [sender]
Done. [sender]
Received: 0 [receiver]
Received: 1 [receiver]
Received: 2 [receiver]
Received: 3 [receiver]
Received: 4 [receiver]
Done. [receiver]

Program exited.
*/
