package main

import "fmt"

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
}
