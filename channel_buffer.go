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
	out := func() {
		x := <-d
		fmt.Println("got", x)

	}
	for i := 1; i < 10; i++ {
		for j := 0; j <= i; j++ {
			go sum(i, j, d)
			go out()
			//fmt.Println(<-d)
		}
	}

}
