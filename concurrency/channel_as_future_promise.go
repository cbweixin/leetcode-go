package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Return receive-only channels as results

// In the following example, the values of two arguments of the sumSquares function call are requested concurrently.
// Each of the two channel receive operations will block until a send operation performs on the corresponding channel.
// It takes about three seconds instead of six seconds to return the final result.
func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(3 * time.Second)
		r <- rand.Int31n(100)
	}()

	return r
}

// Pass send-only channels as arguments

// Same as the last example, in the following example, the values of two arguments of the sumSquares function call are
// requested concurrently. Different to the last example, the longTimeRequest2 function takes a send-only channel as
// parameter instead of returning a receive-only channel result.
func longTimeRequest2(r chan<- int32) {
	time.Sleep(3 * time.Second)
	r <- rand.Int31n(100)

}

// The first response wins

// This is the enhancement of the using-only-one-channel variant in the last example.
//
// Sometimes, a piece of data can be received from several sources to avoid high latencies. For a lot of factors, the
// response durations of these sources may vary much. Even for a specified source, its response durations are also not
// constant. To make the response duration as short as possible, we can send a request to every source in a separated
// goroutine. Only the first response will be used, other slower ones will be discarded.
//
// Note, if there are N sources, the capacity of the communication channel must be at least N-1, to avoid the goroutines
// corresponding the discarded responses being blocked for ever.
func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// sleep 1s/2s/3s
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func sumSqures(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSqures(<-a, <-b))

	c, d := make(chan int32), make(chan int32)
	go longTimeRequest2(c)
	go longTimeRequest2(d)
	fmt.Println(sumSqures(<-c, <-d))

	// instead of using 2 channels c, d, we could use a buffered channel e with cap 2
	e := make(chan int32, 2)
	go longTimeRequest2(e)
	go longTimeRequest2(e)
	fmt.Println(sumSqures(<-e, <-e))

	startTime := time.Now()
	// c2 must be a buffered channel
	c2 := make(chan int32, 5)
	for i := 0; i < cap(c2); i++ {
		go source(c2)
	}

	rnd := <-c2
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
