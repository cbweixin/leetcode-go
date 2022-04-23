package channel_sample

import (
	"fmt"
	"time"
)

func Blocking_Channel() {
	sendingInterval := time.Second
	receptionInterval := time.Second * 2
	intChan := make(chan int, 0)
	go func() {
		var ts0, ts1 int64
		for i := 1; i <= 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent: %d [interval: %d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()
	var ts0, ts1 int64
Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Received:", v)
			} else {
				fmt.Printf("Received: %d [interval: %d s]\n", v, ts1-ts0)
			}
		}
		ts0 = time.Now().Unix()
		time.Sleep(receptionInterval)
	}
	fmt.Println("End.")
}

/**
Sent: 1
Received: 1
Received: 2 [interval: 2 s]
Sent: 2 [interval: 2 s]
Received: 3 [interval: 2 s]
Sent: 3 [interval: 2 s]
Received: 4 [interval: 2 s]
Sent: 4 [interval: 2 s]
Received: 5 [interval: 2 s]
Sent: 5 [interval: 2 s]
End.
happens-before 原则

向此类通道发送元素值的操作会被阻塞，直到至少有一个针对该 通道的接收操作进行为止。该接收操作会先得到元素值的副本，
然后在唤醒发送方所在的goroutine之后返回。也就是说，这时的 接收操作会在对应的发送操作完成之前完成。

从此类通道接收元素值的操作会被阻塞，直到至少有一个针对该 通道的发送操作进行为止。该发送操作会直接把元素值复制给接
收方，然后在唤醒接收方所在的goroutine之后返回。也就是说， 这时的发送操作会在对应的接收操作完成之前完成。
*/
