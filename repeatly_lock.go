package main

import (
	"fmt"
	"sync"
	"time"
)

/*
对于同一个互斥锁的锁定操作和解锁操作总是应该成对的出现。如果我们锁定了一个已被锁定的互斥锁，那么进行重复锁定操作的Goroutine将会被阻塞，
直到该互斥锁回到解锁状态。
*/
func repeatedlyLock() {

	var mutex sync.Mutex
	fmt.Println("Lock the lock. (G0)")
	mutex.Lock()
	fmt.Println("The lock is locked. (G0)")
	/*
		首先，在repeatedlyLock函数被执行伊始，对互斥锁的第一次锁定操作便被进行并顺利地完成。这由第一行和第二行打印内容可以看出。而后，
		在repeatedlyLock函数中被启用的那三个Goroutine在G0的第一次“睡眠”期间开始被运行。当相应的go函数中的对互斥锁的锁定操作被进行的时候，
		它们都被阻塞住了。原因是该互斥锁已处于锁定状态了。这就是我们在这里只看到了三个连续的Lock the lock. (G<i>)
		而没有立即看到The lock is locked. (G<i>)的原因。随后，G0“睡醒”并解锁互斥锁。这使得正在被阻塞的G1、G2和G3都会有机会重新锁定该互斥锁。
		但是，只有一个Goroutine会成功。成功完成锁定操作的某一个Goroutine会继续执行在该操作之后的语句。而其他Goroutine将继续被阻塞，
		直到有新的机会到来。这也就是上述打印内容中的最后三行所表达的含义。
	*/
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (G%d)\n", i)

		}(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Unlock the lock. (G0)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (G0)")
	time.Sleep(time.Second)
}

func main() {
	repeatedlyLock()

}
