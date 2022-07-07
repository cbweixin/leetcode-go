package main

import (
	"fmt"
	"time"
)

// 只有生产者，没有消费者
func f1() {
	ch := make(chan int)
	ch <- 1
}

// 只有消费者，没有生产者
func f2() {
	ch := make(chan int)
	<-ch
}

// 生产者和消费者出现在同一个 goroutine 中
func f3() {
	ch := make(chan int)
	// producer blocked the routine, consumer has no opportunity to run
	ch <- 1
	// this line not able to reach
	<-ch

}

func leak1() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 100
	}()

	select {
	case <-time.After(2500 * time.Millisecond):
		fmt.Println("timeout! exit...")
	case result := <-ch:
		fmt.Printf("result : %n\n", result)
	}
}

func leak2() {
	ch := make(chan int)

	// consumer
	go func() {
		for result := range ch {
			fmt.Printf("result : %d\n", result)
		}
	}()

	// producer
	ch <- 1
	ch <- 2
	time.Sleep(time.Second)
	fmt.Println("producer close channel")
	close(ch)

}

func panic1() {
	// panic, try to send to a closed channel
	ch := make(chan int, 1)
	close(ch)
	ch <- 1
}

func panic2() {
	ch := make(chan int, 1)
	done := make(chan struct{}, 1)

	go func() {
		<-time.After(2 * time.Second)
		fmt.Println("close channel 2")
		close(ch)
		close(done)
	}()

	go func() {
		<-time.After(1 * time.Second)
		fmt.Println("close 1")
		ch <- 1
		close(ch)
	}()
	<-done

}

func main() {
	panic2()

	// f2()
	// f1()
	// f3()
	// fmt.Printf("before goroutine num: %d\n", runtime.NumGoroutine())
	// for i := 0; i < 10; i++ {
	// 	leak2()
	// }
	// time.Sleep(time.Second)
	// fmt.Printf("after goroutine num: %d\n", runtime.NumGoroutine())
	//
	// print("hello")
}
