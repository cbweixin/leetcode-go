package main

import (
	"fmt"
	"runtime"
)

type ZeroEvenOdd2 struct {
	n                int
	streamEvenToZero chan struct{}
	streamOddToZero  chan struct{}
	streamZeroToEven chan struct{}
	streamZeroToOdd  chan struct{}
	streamZeroToEnd  chan struct{}
}

func (self *ZeroEvenOdd2) Zero(zero func(int)) {
	for i := 0; i < self.n; i++ {
		select {
		case <-self.streamOddToZero:
			zero(0)
			self.streamZeroToEven <- struct{}{}
		case <-self.streamEvenToZero:
			zero(0)
			self.streamZeroToOdd <- struct{}{}
		default: // what does this mean?
			runtime.Gosched()
			i--
		}
	}

	if self.n%2 == 0 {
		<-self.streamEvenToZero
	} else {
		<-self.streamOddToZero
	}

	self.streamZeroToEnd <- struct{}{}

}

func (self *ZeroEvenOdd2) Odd(odd func(int)) {

	for i := 1; i <= self.n; i += 2 {
		<-self.streamZeroToOdd
		odd(i)
		self.streamOddToZero <- struct{}{}
	}

}

func (self *ZeroEvenOdd2) Even(even func(int)) {

	for i := 2; i <= self.n; i += 2 {
		<-self.streamZeroToEven
		even(i)
		self.streamEvenToZero <- struct{}{}
	}

}

func PrintNumber(x int) {
	fmt.Printf("%d", x)
}

func main() {
	var PrintZeroEvenOdd = func(testNum int) {
		var zeo = &ZeroEvenOdd2{
			n:                testNum,
			streamEvenToZero: make(chan struct{}),
			streamOddToZero:  make(chan struct{}),
			streamZeroToEven: make(chan struct{}),
			streamZeroToOdd:  make(chan struct{}),
			streamZeroToEnd:  make(chan struct{}),
		}

		//go func() { zeo.streamEvenToZero <- struct{}{} }()
		go zeo.Zero(PrintNumber)
		go zeo.Even(PrintNumber)
		go zeo.Odd(PrintNumber)
		zeo.streamEvenToZero <- struct{}{}

		<-zeo.streamZeroToEnd
		fmt.Println()
	}

	for testNum := range [14]int{} {
		fmt.Printf("Case %2d: ", testNum)
		PrintZeroEvenOdd(testNum)
	}

}
