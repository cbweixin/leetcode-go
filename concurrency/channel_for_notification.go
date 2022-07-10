package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

// Use Channels for Notifications

// Notifications can be viewed as special requests/responses in which the responded values are not important. Generally
// , we use the blank struct type struct{} as the element types of the notification channels, for the size of type
// struct{} is zero, hence values of struct{} doesn't consume memory.

type T struct{}

// N-to-1 and 1-to-N notifications

// By extending the above two use cases a little, it is easy to do N-to-1 and 1-to-N notifications.
// In fact, the ways to do 1-to-N and N-to-1 notifications introduced in this sub-section are not used commonly in
// practice. In practice, we often use sync.WaitGroup to do N-to-1 notifications, and we do 1-to-N notifications by
// close channels. Please read the next sub-section for details.

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait for notification

	log.Print("worker #", id, " starts")
	// simulate a work load
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("worker #", id, " job done")

	done <- T{}
}

func main() {
	// 1-to-1 notification by sending a value to a channel

	// If there are no values to be received from a channel, then the next receive operation on the channel will block
	// until another goroutine sends a value to the channel. So we can send a value to a channel to notify another
	// goroutine which is waiting to receive a value from the same channel.
	//
	// In the following example, the channel done is used as a signal channel to do notifications.

	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // can be buffered or not

	// sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()

	<-done
	//do other things

	fmt.Println(values[0], len(values), values[len(values)-1])

	// 1-to-1 notification by receiving a value from a channel

	// If the value buffer queue of a channel is full (the buffer queue of an unbuffered channel is always full), a
	// send operation on the channel will block until another goroutine receives a value from the channel. So we can
	// receive a value from a channel to notify another goroutine which is waiting to send a value to the same channel.
	// Generally, the channel should be an unbuffered channel.
	//
	// This notification way is used much less common than the way introduced in the last example.

	done1 := make(chan struct{})
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		<-done1
	}()

	done1 <- struct{}{}
	fmt.Println(values[0], len(values), values[len(values)-1])

	log.SetFlags(0)
	ready, done3 := make(chan T), make(chan T)
	go worker(0, ready, done3)
	go worker(1, ready, done3)
	go worker(2, ready, done3)

	// simulate an initialization phrase
	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications
	ready <- T{}
	ready <- T{}
	ready <- T{}

	// being N-to-1 notified
	<-done3
	<-done3
	<-done3

	ready1, done4 := make(chan T), make(chan T)
	go worker(3, ready1, done4)
	go worker(4, ready1, done4)
	go worker(5, ready1, done4)

	// simulate an initialization phrase
	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications

	// Broadcast (1-to-N) notifications by closing a channel
	// The way to do 1-to-N notifications shown in the last sub-section is seldom used in practice, for there is a
	// better way. By making using of the feature that infinite values can be received from a closed channel, we can
	// close a channel to broadcast notifications.
	//
	// By the example in the last sub-section, we can replace the three channel send operations ready <- struct{}{} in
	// the last example with one channel close operation close(ready) to do an 1-to-N notifications.
	close(ready1) // broadcast notification

	// being N-to-1 notified
	<-done4
	<-done4
	<-done4
}
