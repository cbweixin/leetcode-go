package main

import (
	"fmt"
	"math/rand"
	"runtime"
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

// Another way to implement the first-response-wins use case

// As mentioned above, we can use the select mechanism (try-send) with a buffered channel which capacity is one
// (at least) to implement the first-response-wins use case. For example,
func source2(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	fmt.Println("source2 before branch ", ra)
	time.Sleep(time.Second * time.Duration(rb))
	select {
	case c <- ra:
	default:
		fmt.Println("source2 default branch ", ra)
	}

}

// More request-response variants
// The parameter and result channels can be buffered so that the response sides won't need to wait for the request sides
// to take out the transferred values.
//
// Sometimes, a request is not guaranteed to be responded back a valid value. For all kinds of reasons, an error may be
// returned instead. For such cases, we can use a struct type like struct{v T; err error} or a blank interface type as
// the channel element type.
//
// Sometimes, for some reasons, the response may need a much longer time than the expected to arrive, or will never
// arrive. We can use the timeout mechanism introduced below to handle such circumstances.
//
// Sometimes, a sequence of values may be returned from the response side, this is kind of the data flow mechanism
// mentioned later below.

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
	// if c2 is not a buffered channel, then c <- ra would be a blocking call, then goroutine would stucked there
	// and cause memory leak
	c2 := make(chan int32, 5)
	for i := 0; i < cap(c2); i++ {
		go source(c2)
	}

	rnd := <-c2
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)

	time.Sleep(time.Second * 2)

	// c3 must be a buffered channel, reason is different from above.
	// if c3 is not a buffered channel, then c3 <- ra would be a blocking call, hence it would go to default branch
	// thus, rnd := <- c3 might never be getting any value
	fmt.Println("before ", runtime.NumGoroutine())
	c3 := make(chan int32, 0)
	for i := 0; i < 5; i++ {
		go source2(c3)
	}

	fmt.Println("middle ", runtime.NumGoroutine())
	time.Sleep(time.Second * 1)
	rnd2 := <-c3
	println(rnd2)
	time.Sleep(time.Second * 1)
	fmt.Println("after ", runtime.NumGoroutine())

}
