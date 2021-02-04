package main

import "fmt"

//Suppose we have a class:
//
//
//public class Foo {
//  public void first() { print("first"); }
//  public void second() { print("second"); }
//  public void third() { print("third"); }
//}
//
//
// The same instance of Foo will be passed to three different threads. Thread A
//will call first(), thread B will call second(), and thread C will call third().
//Design a mechanism and modify the program to ensure that second() is executed af
//ter first(), and third() is executed after second().
//
//
//
// Example 1:
//
//
//Input: [1,2,3]
//Output: "firstsecondthird"
//Explanation: There are three threads being fired asynchronously. The input [1,
//2,3] means thread A calls first(), thread B calls second(), and thread C calls t
//hird(). "firstsecondthird" is the correct output.
//
//
// Example 2:
//
//
//Input: [1,3,2]
//Output: "firstsecondthird"
//Explanation: The input [1,3,2] means thread A calls first(), thread B calls th
//ird(), and thread C calls second(). "firstsecondthird" is the correct output.
//
//
//
// Note:
//
// We do not know how the threads will be scheduled in the operating system, eve
//n though the numbers in the input seems to imply the ordering. The input format
//you see is mainly to ensure our tests' comprehensiveness.
// 👍 659 👎 114

// 2021-02-03 17:20:03

//There is no code of Go type for this problem

/*本题考核点？
指定各种不同顺序执行 First(), Second(), Third() 三个 goroutine，但三者都必须以不变顺序印出字串，印出顺序不受顺序执行影响。
goroutine 若不刻意控制，将无法保证执行的先后顺序，因此本题就是要考核对 goroutine 顺序控制的能力。
https://zhuanlan.zhihu.com/p/103405131
*/

func First(streamSync [3]chan interface{}) {
	fmt.Print("First")
	streamSync[0] <- nil
}

func Second(streamSync [3]chan interface{}) {
	<-streamSync[0]
	fmt.Print("Second")
	streamSync[1] <- nil

}

func Third(streamSync [3]chan interface{}) {
	<-streamSync[1]
	fmt.Print("Third")
	streamSync[2] <- nil
}

func PrintInOrder(callOrder [3]int) {
	order := callOrder
	fmt.Println("[] call order:", order)

	var streamSync [3]chan interface{}
	for i := range streamSync {
		streamSync[i] = make(chan interface{})
	}
	var functionNumTable = map[int]func([3]chan interface{}){
		1: First,
		2: Second,
		3: Third,
	}

	for _, fNum := range order {
		go functionNumTable[fNum](streamSync)
	}

	<-streamSync[2]

}

func main() {
	var testCases = [][3]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, theCase := range testCases {
		PrintInOrder(theCase)
		fmt.Println()
		fmt.Println()
	}
}
