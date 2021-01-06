package main

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f_f2(left, right chan int) { left <- 1 + <-right }

func main() {
	// 译者注：原本认为leftmost的结果为1，认为只在最初做了一次赋值，实际结果为100000（无缓存信道具有同步阻塞的特性）
	//
	//1.主线程的right <- 0，right不是最初循环的那个right，而是最终循环的right
	//
	//2.for循环中最初的go f(left, right)因为没有发送者一直处于等待状态
	//
	//3.当主线程的right <- 0执行时，类似于递归函数在最内层产生返回值一般
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f_f2(left, right)
	}
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, ongeveer 1,5 s

}
