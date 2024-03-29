package main

import (
	"fmt"
	"time"
)
import "./ratelimiter"

// Try-Send and Try-Receive

// A select block with one default branch and only one case branch is called a try-send or try-receive channel operation
// , depending on whether the channel operation following the case keyword is a channel send or receive operation.
// If the operation following the case keyword is a send operation, then the select block is called as try-send
// operation. If the send operation would block, then the default branch will get executed (fail to send), otherwise,
// the send succeeds and the only case branch will get executed.
// If the operation following the case keyword is a receive operation, then the select block is called as try-receive
// operation. If the receive operation would block, then the default branch will get executed (fail to receive),
// otherwise, the receive succeeds and the only case branch will get executed.
// Try-send and try-receive operations never block.
//
// The standard Go compiler makes special optimizations for try-send and try-receive select blocks, their execution
// efficiencies are much higher than multi-case select blocks.
//
// The following is an example which shows how try-send and try-receive work.

func main() {
	type Book struct {
		id int
	}
	bookShelf := make(chan Book, 3)

	// try-send
	for i := 0; i < cap(bookShelf)*2; i++ {
		select {
		case bookShelf <- Book{id: i}:
			fmt.Println("success to put book ", i)
		default:
			fmt.Println("failed to put book")
		}
	}

	// try-receive
	for i := 0; i < cap(bookShelf)*2; i++ {
		select {
		case book := <-bookShelf:
			fmt.Println("succeeded to get book ", book.id)
		default:
			fmt.Println("failed to get book")
		}
	}

	requests := make(chan ratelimiter.Request)
	go ratelimiter.HandleRequests(requests)
	// with below sleep, the rate limiter would gather 1 minutes ticker, once wake up
	// the below for loop would quickly be processed , we will see 1-200 got bursted print out
	// then 201, 202, each one got printed individually
	time.Sleep(time.Minute)

	for i := 0; ; i++ {
		requests <- i
	}

}
