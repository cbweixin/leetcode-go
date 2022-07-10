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

type Seat int
type Bar chan Seat

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait for notification

	log.Print("worker #", id, " starts")
	// simulate a work load
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("worker #", id, " job done")

	done <- T{}
}

// Timer: scheduled notification
// n fact, the After function in the time standard package provides the same functionality, with a much more
// efficient implementation. We should use that function instead to make the code look clean.
//
// Please note, <-time.After(aDuration) will make the current goroutine enter blocking state, but a
// time.Sleep(aDuration) function call will not.
//
// The use of <-time.After(aDuration) is often used in the timeout mechanism which will be introduced below.
func AfterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()

	return c
}

// Use Channels as Mutex Locks

// One of the above examples has mentioned that one-capacity buffered channels can be used as one-time binary semaphore.
// In fact, such channels can also be used as multi-time binary semaphores, a.k.a., mutex locks, though such mutex locks
// are not efficient as the mutexes provided in the sync standard package.
//
// There are two manners to use one-capacity buffered channels as mutex locks.
// Lock through a send, unlock through a receive.
// Lock through a receive, unlock through a send.
// The following is a lock-through-send example.
func mutexEx1() {
	mutex := make(chan T, 1)
	counter := 0

	increase := func() {
		mutex <- T{}
		counter++
		<-mutex
	}

	increase1000 := func(done chan<- T) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- T{}
	}

	done := make(chan T)
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println("counter is ", counter)
}

// The following is a lock-through-receive example. It just shows the modified part based on the above lock-through-send
// example.
func mutexEx2() {
	mutex := make(chan T, 1)
	counter := 0
	mutex <- T{}

	increase := func() {
		<-mutex
		counter++
		mutex <- T{}
	}

	increase1000 := func(done chan<- T) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- T{}
	}

	done := make(chan T)
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println("counter is ", counter)
}

// Use Channels as Counting Semaphores

// Buffered channels can be used as counting semaphores. Counting semaphores can be viewed as multi-owner locks. If the
// capacity of a channel is N, then it can be viewed as a lock which can have most N owners at any time. Binary
// semaphores (mutexes) are special counting semaphores, each of binary semaphores can have at most one owner at
// any time.
//
// Counting semaphores are often used to enforce a maximum number of concurrent requests.
//
// Like using channels as mutexes, there are also two manners to acquire one piece of ownership of a channel semaphore.
// Acquire ownership through a send, release through a receive.
// Acquire ownership through a receive, release through a send.
// An example of acquiring ownership through receiving values from a channel.
func (bar Bar) ServeCustomer(c int) {
	log.Print("customer #", c, " enters the bar")
	seat := <-bar // need a seat to drink
	log.Print("++ customer #", c, " drinks at seat #", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func runBar() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // the bar has 10 seats
	// place seat in bar
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// none of the sends will block
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}

	// The last for loop in the main function is to avoid the program exiting. There is a better way, which will be
	// introduced below, to do the job.
	for { // sleeping != blocking
		time.Sleep(time.Second)
	}

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
	// Surely, we can also close a channel to do a 1-to-1 notification. In fact, this is the most used notification way
	// in Go.
	//
	// The feature that infinite values can be received from a closed channel will be utilized in many other use cases
	// introduced below. In fact, the feature is used popularly in the standard packages. For example, the context
	// package uses this feature to confirm cancellations.
	close(ready1) // broadcast notification

	// being N-to-1 notified
	<-done4
	<-done4
	<-done4

	fmt.Println("Hi!")
	<-AfterDuration(time.Second)
	fmt.Println("Hello!")
	<-AfterDuration(time.Second)
	fmt.Println("Bye!")

	mutexEx1()
	mutexEx2()

	runBar()
}
