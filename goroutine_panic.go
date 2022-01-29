package main

import (
	"fmt"
	"os"
)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch) // if this is ommitted: panic: all goroutines are asleep - deadlock!
}

func tel2(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func main() {
	var ok = true
	var i int
	ch := make(chan int)

	go tel(ch)
	for ok {
		if i, ok = <-ch; ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		}
	}

	var ok2 = true
	quit := make(chan bool)
	ch2 := make(chan int)
	go tel2(ch2, quit)

	for ok2 {
		select {
		case i := <-ch2:
			fmt.Printf("the counter is %d\n", i)

		case <-quit:
			os.Exit(0)
		}
	}

	// time.Sleep(1e9)

}
