package main

import (
	"fmt"
	"time"
)

//Suppose you are given the following code:
//
//
//class FooBar {
//  public void foo() {
//    for (int i = 0; i < n; i++) {
//      print("foo");
//    }
//  }
//
//  public void bar() {
//    for (int i = 0; i < n; i++) {
//      print("bar");
//    }
//  }
//}
//
//
// The same instance of FooBar will be passed to two different threads. Thread A
// will call foo() while thread B will call bar(). Modify the given program to out
//put "foobar" n times.
//
//
//
// Example 1:
//
//
//Input: n = 1
//Output: "foobar"
//Explanation: There are two threads being fired asynchronously. One of them cal
//ls foo(), while the other calls bar(). "foobar" is being output 1 time.
//
//
// Example 2:
//
//
//Input: n = 2
//Output: "foobarfoobar"
//Explanation: "foobar" is being output 2 times.
//
// 👍 335 👎 29

// 2021-02-03 17:43:37

//There is no code of Go type for this problem

type FooBar struct {
	n              int
	streamFooToBar chan struct{}
	streamBarToFoo chan struct{}
	streamEnd      chan struct{}
}

/**
3. Foo() 与 Bar() 如何收尾？
这里要特别注意的是，Foo() 與 Bar() 只有差异在最后一行，用意是什么？

多个 goroutine 用 unbuffered channel 互相交接棒，会有一个尴尬的情况，就是互为消费者、又互为生产者，因此先 return 的 goroutine 没事，
但是后 return 的 goroutine 会由于消费者消失，send to channel 的时候发生 Deadlock。

根据本题游戏规则， printBar() 一定要比 printFoo() 晚执行，因此不做特别处理的话，会在 Bar() 试图做最后一次交棒时，由于消费者消失发生
Deadlock。

我的应对方式，就是让 Foo() return 前做一次无特别作用的接棒，这样就可以避免 Bar() return 前找不到消费者的问题。

4. 自循环启动
fooBar.streamBarToFoo <- struct{}{} //启动
前面说过，本题解法采用 Foo() 与 Bar() 彼此循环交棒的方式，因此不能自行启动，需要外界给讯号，所以在 main() 执行各个 goroutine 以后以
fooBar.streamBarToFoo <- struct{}{} ，让 main() 假装自己是 Bar() 交棒给 Foo()，以启动交接棒循环。

5. 特别条件下，用 unbuffered channel 取代 sync.WaitGroup
<-fooBar.streamEnd                  //as wg.Wait()
为了等待多个 goroutine 都结束再往下，一般来说用 sync.WaitGroup.Wait() 是标准做法。但这一题还有更轻量的方法。

虽然这一题是 Concurrency，但是各个 goroutine 的结束顺序已经被定死，我们很清楚知道谁负责收尾，所以让负责收尾的 goroutine send to
unbuffered channel，然后在 main() read 掉，这样就不需要使用 sync.WaitGroup。

执行各个 goroutine 以后，main() 会由于 <-fooBar.streamEnd 还没有被传入而被卡住，这就相当于 sync.WaitGroup.Wait() 的作用了。

由于是 Bar() 会做最后一次有意义的执行，因此让 Bar() return 之前执行 this.streamEnd <- struct{}{}，这就相当于交棒给 main()，
main() 终于可以从被卡住的 <-fooBar.streamEnd 往下（因为终于有东西可以读），就像便秘了三天突然畅通一样！

*/

func (this *FooBar) Foo(printFoo func()) {
	for i := 0; i < this.n; {
		<-this.streamBarToFoo
		printFoo()
		i++
		this.streamFooToBar <- struct{}{}
	}

	<-this.streamBarToFoo

}

func (this *FooBar) Bar(printBar func()) {
	for i := 0; i < this.n; {
		<-this.streamFooToBar
		printBar()
		i++
		this.streamBarToFoo <- struct{}{}

	}
	this.streamEnd <- struct{}{}
}

// implementation 2
func foo(n int, oddChan chan bool, evenChan chan bool) {
	for i := 0; i < n; i++ {
		fmt.Print("foo")
		evenChan <- true
		<-oddChan
	}
}

func bar(n int, oddChan chan bool, evenChan chan bool) {
	for i := 0; i < n; i++ {
		<-evenChan
		fmt.Print("bar")
		oddChan <- true

	}
}

func fooBar(n int) {
	var oddChan = make(chan bool)
	var evenChan = make(chan bool)
	go foo(n, oddChan, evenChan)
	go bar(n, oddChan, evenChan)
	time.Sleep(2 * time.Second)
}

func main() {
	var PrintFooBar = func(times int) {

		fooBar := &FooBar{
			n:              times,
			streamFooToBar: make(chan struct{}),
			streamBarToFoo: make(chan struct{}),
			streamEnd:      make(chan struct{}),
		}

		go fooBar.Foo(func() { fmt.Printf("Foo") })
		go fooBar.Bar(func() { fmt.Printf("Bar ") })
		fooBar.streamBarToFoo <- struct{}{} //start
		<-fooBar.streamEnd                  //as wg.Wait()
		fmt.Println()
	}

	testCase := []int{0, 1, 2, 3, 4, 5, 7}

	for _, repeat := range testCase {
		fmt.Printf("Repeat %d: ", repeat)
		PrintFooBar(repeat)
	}

	fooBar(5)
	fmt.Println()
	fooBar(10)
}
