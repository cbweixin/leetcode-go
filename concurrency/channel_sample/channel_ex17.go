package channel_sample

import (
	"fmt"
	"time"
)

func Ticker_test() {
	// here we use a buffered channel. if we use unbuffered channel, it is possible
	// all case are blocking call, so the whole goroutine just stucked
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Println("End. [sender]")
	}()
	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			break
		}
	}

	/**
	Received: 1
	Received: 2
	Received: 3
	Received: 1
	Received: 2
	Received: 3
	Got: 12
	End. [receiver]
	1,2,3 is randomly generated, but no matter stop ticker or not, I will always see "End. [receiver]"
	*/
	ticker.Stop()
	fmt.Println("End. [receiver]")
}

/**
Received: 1
Received: 3
Received: 3
Received: 1
Received: 3
Got: 11
End. [receiver]
... OR
Received: 2
Received: 2
Received: 2
Received: 1
Received: 1
Received: 1
Received: 2
Got: 11
End. [receiver]

*/
