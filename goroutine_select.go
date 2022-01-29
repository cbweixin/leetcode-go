package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump4(ch1)
	go pump5(ch2)
	go suck4(ch1, ch2)

	time.Sleep(1e5)

}

func pump4(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump5(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck4(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
