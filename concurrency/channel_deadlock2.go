package main

import "time"

// https://stackoverflow.com/questions/5983544/go-channels-and-deadlock
func main() {
	c1 := make(chan int, 1)
	c2 := make(chan int)

	go func() {
		for i := range c1 {
			println("G1 got", i)
			c2 <- i
		}
	}()

	go func() {
		for i := range c2 {
			println("G2 got", i)
			c1 <- i
		}
	}()

	c1 <- 1

	time.Sleep(time.Millisecond * 1)

	// in fact this doesn't cause deadlock at 1.17 and 1.18 version
	// the above stackoverflow questions asked 11 years ago, during this period of time, it evloved
	c1 <- 2

	time.Sleep(time.Millisecond * 2)
}
