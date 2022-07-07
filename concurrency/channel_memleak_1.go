package main

import (
	"fmt"
	"runtime"
	"time"
)

func sleep(key string, sec int) {
	fmt.Printf("before sleep %d seconds: %s.\n", sec, key)
	time.Sleep(time.Duration(sec) * time.Second) // sleep i 秒，time.Sleep()
	fmt.Printf("after sleep %d seconds: %s.\n", sec, key)
}

func read(ch chan int) {
	fmt.Printf("before read from channel: %v\n", ch)
	var val int = <-ch
	fmt.Printf("after read from channel: %d %v\n", val, ch)
}

func write(ch chan int, val int) {
	fmt.Printf("before write to channel: %d %v\n", val, ch)
	ch <- val
	fmt.Printf("after write to channel: %d %v\n", val, ch)
}

func TestChan3() {
	fmt.Println("@@@@@@@@@@@@ TestChan 3")

	// cpu num: 16
	fmt.Printf("cpu num: %d\n", runtime.NumCPU())

	// 虽然goroutine是并发执行的，但是它们并不是并行运行的。如果不告诉Go额外的东西，同
	// 一时刻只会有一个goroutine执行。利用runtime.GOMAXPROCS(n)可以设置goroutine
	// 并行执行的数量。GOMAXPROCS 设置了同时运行的CPU 的最大数量，并返回之前的设置。
	val := runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	// last goroutine num: 16
	fmt.Printf("last goroutine num: %d\n", val)

	// goroutine num: 1
	fmt.Printf("goroutine num: %d\n", runtime.NumGoroutine())

	var ch1 chan int = make(chan int, 0)
	var ch2 chan int = make(chan int, 0)
	var ch3 chan int = make(chan int, 0)

	go write(ch1, 22)
	go write(ch2, 33)
	go write(ch3, 44)
	go read(ch1)
	go read(ch2)
	go read(ch3)

	// goroutine num: 7
	fmt.Printf("goroutine num: %d\n", runtime.NumGoroutine())
	sleep("TestChan3", 3)
}

func main() {
	TestChan3()
}
