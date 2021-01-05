package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump2()
	go suck2(stream)
	time.Sleep(1e1)
	fmt.Println("-------hi----------------------------")
	suck3(pump3())
	time.Sleep(1e1)

}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck2(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck3(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
