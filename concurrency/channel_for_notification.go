package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

// Use Channels for Notifications

// Notifications can be viewed as special requests/responses in which the responded values are not important. Generally
// , we use the blank struct type struct{} as the element types of the notification channels, for the size of type
// struct{} is zero, hence values of struct{} doesn't consume memory.

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
}
