package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
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

	select { // if timeout too early, then above goroutine would blocking forever,cause memory leak
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
		// for..range would block consumer, if producer do not close channel, then this goroutine would stucked
		for result := range ch {
			fmt.Printf("result : %d\n", result)
		}
	}()

	// producer
	ch <- 1
	ch <- 2
	time.Sleep(time.Second)
	fmt.Println("producer close channel")
	// without this line, OOM might occur
	// https://cbweixin.github.io/2022/07/06/channel-related-issues/
	close(ch)

}

func panic1() {
	// panic, try to send to a closed channel
	// but try to get from a closed channel, you would get zero value, no panic would happen
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
	// panic2()

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

	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh 是额外引入的一个信号 channel.
	// 它的生产者是下面的 toStop channel，
	// 消费者是上面 dataCh 的生产者和消费者
	toStop := make(chan string, 1)
	// toStop 是拿来关闭 stopCh 用的，由 dataCh 的生产者和消费者写入
	// 由下面的匿名中介函数(moderator)消费
	// 要注意，这个一定要是 buffered channel （否则没法用 try-send 来处理了）
	var stoppedBy string // moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()
	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// try-send 操作
					// 如果 toStop 满了，就会走 default 分支啥也不干，也不会阻塞
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}
				// try-receive 操作，尽快退出
				// 如果没有这一步，下面的 select 操作可能造成 panic
				select {
				case <-stopCh:
					return
				default:
				}
				// 如果尝试从 stopCh 取数据的同时，也尝试向 dataCh
				// 写数据，则会命中 select 的伪随机逻辑，可能会写入数据
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}
	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			for {
				// 同上
				select {
				case <-stopCh:
					return
				default:
				}
				// 尝试读数据
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)

}
