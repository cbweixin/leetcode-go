package main

import "fmt"

//Suppose you are given the following code:
//
//
//class ZeroEvenOdd {
//  public ZeroEvenOdd(int n) { ... }      // constructor
//  public void zero(printNumber) { ... }  // only output 0's
//  public void even(printNumber) { ... }  // only output even numbers
//  public void odd(printNumber) { ... }   // only output odd numbers
//}
//
//
// The same instance of ZeroEvenOdd will be passed to three different threads:
//
//
// Thread A will call zero() which should only output 0's.
// Thread B will call even() which should only ouput even numbers.
// Thread C will call odd() which should only output odd numbers.
//
//
// Each of the threads is given a printNumber method to output an integer. Modif
//y the given program to output the series 010203040506... where the length of the
// series must be 2n.
//
//
//
// Example 1:
//
//
//Input: n = 2
//Output: "0102"
//Explanation: There are three threads being fired asynchronously. One of them c
//alls zero(), the other calls even(), and the last one calls odd(). "0102" is the
// correct output.
//
//
// Example 2:
//
//
//Input: n = 5
//Output: "0102030405"
//
// 👍 221 👎 149

// 2021-02-03 20:37:50

//There is no code of Go type for this problem

type ZeroEvenOdd struct {
	n          int
	zeroStream chan struct{}
	evenStream chan struct{}
	oddStream  chan struct{}
	endStream  chan struct{}
}

func (this *ZeroEvenOdd) Zero(zero func(n int)) {
	for i := 0; i < this.n; {
		<-this.zeroStream
		zero(0)
		i++
		if i%2 == 0 {
			this.evenStream <- struct{}{}
		} else {
			this.oddStream <- struct{}{}
		}
	}

	<-this.zeroStream
}

func (this *ZeroEvenOdd) Even(even func(n int)) {
	i := 0
	for i < this.n {
		<-this.evenStream
		even(i)
		i += 2
		this.zeroStream <- struct{}{}

	}
	if i >= this.n {
		this.zeroStream <- struct{}{}
		this.endStream <- struct{}{}
	}
}

func (this *ZeroEvenOdd) Odd(odd func(n int)) {
	i := 1
	for i < this.n {
		<-this.evenStream
		odd(i)
		i += 2
		this.zeroStream <- struct{}{}

	}
	if i >= this.n {
		this.zeroStream <- struct{}{}
		this.endStream <- struct{}{}
	}
}

func main() {
	var printZeo = func(num int) {

		zeo := &ZeroEvenOdd{
			n:          num,
			zeroStream: make(chan struct{}),
			evenStream: make(chan struct{}),
			oddStream:  make(chan struct{}),
			endStream:  make(chan struct{}),
		}

		go zeo.Zero(func(n int) { fmt.Print(n) })
		go zeo.Odd(func(n int) { fmt.Print(n) })
		go zeo.Even(func(n int) { fmt.Print(n) })

		zeo.zeroStream <- struct{}{}
		<-zeo.endStream
		fmt.Println()

	}

	/*	testCase := []int{1, 2, 3, 4, 5, 7,8,9,10}

		for _, repeat := range testCase {
			fmt.Printf("Repeat %d: ", repeat)
			printZeo(repeat)
		}
	*/

	printZeo(1)
}
