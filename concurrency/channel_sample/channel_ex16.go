package channel_sample

import (
	"fmt"
	"time"
)

func Timer_test3() {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()
	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else { // re-use timer
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Printf("Received: %v\n", e)
		case <-timer.C:
			fmt.Println("Timeout!")
		}
	}

}

/**
Timeout!
Received: 0
Timeout!
Timeout!
Timeout!
Received: 1
Timeout!
Timeout!
Received: 2
Timeout!
Timeout!
Received: 3
Timeout!
Timeout!
Received: 4
End.
*/
