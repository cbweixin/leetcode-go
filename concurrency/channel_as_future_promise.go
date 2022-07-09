package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Return receive-only channels as results
func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(3 * time.Second)
		r <- rand.Int31n(100)
	}()

	return r
}

// Pass send-only channels as arguments
func longTimeRequest2(r chan<- int32) {
	time.Sleep(3 * time.Second)
	r <- rand.Int31n(100)

}
func sumSqures(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSqures(<-a, <-b))

	c, d := make(chan int32), make(chan int32)
	longTimeRequest2(c)
	longTimeRequest2(d)
	fmt.Println(<-c, <-d)
}
