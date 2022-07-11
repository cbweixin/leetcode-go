package main

import "fmt"

// Channel Encapsulated in Channel

// Sometimes, we can use a channel type as the element type of another channel type. In the following example,
// chan chan<- int is a channel type which element type is a send-only channel type chan<- int.

var counter = func(n int) chan<- chan<- int {
	requests := make(chan chan<- int)

	go func() {
		for request := range requests {
			if request == nil {
				n++
			} else {
				request <- n
			}
		}
	}()

	return requests

}(0)

func main() {
	increaese1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			// counter return requests : chan chan<- int
			// nil has been put into rqeuests chan for 1000 times
			counter <- nil
		}
		// blocking here, until receiver read value from done
		done <- struct{}{}
	}

	done := make(chan struct{})

	go increaese1000(done)
	go increaese1000(done)
	<-done
	<-done

	request := make(chan int, 1)
	counter <- request

	fmt.Println(<-request)

}
