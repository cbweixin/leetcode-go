package main

import (
	"fmt"
	"time"
)

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	c := make(chan int, 50)
	go func() {
		time.Sleep(1 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
	time.Sleep(2 * 1e9)

	d := make(chan int)
	go sum(12, 13, d)
	fmt.Println(<-d) // 25

}
