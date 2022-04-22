package channel_sample

import "fmt"

func Select_test3() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
	Loop: // notice the usage of label
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("End.")
					// be aware of the tag LOOP , without it, you just break `select` instead of `for` loop
					break Loop // how to break a label
				}
				fmt.Printf("Received: %v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}

/**
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
Received: 6
Received: 7
Received: 8
Received: 9
End.
*/
