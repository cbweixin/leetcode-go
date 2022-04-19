package channel_sample

import "fmt"

func Chan_Conv() {
	var ok bool
	ch := make(chan int, 1)
	// true
	_, ok = interface{}(ch).(chan int)
	fmt.Println("chan int => chan int:", ok)
	// false, channel with different direction is different type
	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int:", ok)
	// false
	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("chan int => chan<- int:", ok)

	sch := make(chan<- int, 1)
	// true
	_, ok = interface{}(sch).(chan<- int)
	fmt.Println("chan<- int => chan<- int:", ok)
	// false
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan<- int => chan int:", ok)

	rch := make(<-chan int, 1)
	// true
	_, ok = interface{}(rch).(<-chan int)
	fmt.Println("<-chan int => <-chan int:", ok)
	// false
	_, ok = interface{}(rch).(chan int)
	fmt.Println("<-chan int => chan int:", ok)
}

/**
chan int => chan int: true
chan int => <-chan int: false
chan int => chan<- int: false
chan<- int => chan<- int: true
chan<- int => chan int: false
<-chan int => <-chan int: true
<-chan int => chan int: false

Program exited.
*/
