package main

import (
	"fmt"
	"time"
)

// review : d:\record\go_routine_fibnacci.mp4
func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			//fmt.Printf("x is %d\n", x)
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		//fmt.Printf("t is %d\n",t)
		for {
			x <- <-a + <-b
		}
	}()
	t := <-out
	fmt.Printf("t again is %d\n", t)
	return out
}

func main() {

	start := time.Now()
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}
