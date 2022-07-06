package main

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

func main() {
	go f2()
	go f1()
	f3()
	print("hello")
}
