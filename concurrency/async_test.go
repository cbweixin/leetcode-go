package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/sun0day/async"
)

func main() {
	a := 1
	b := 2
	f := func() int {
		c := a + b
		fmt.Printf("async result=%d\n", c)
		return c
	}

	af := async.Async[int](f)

	fmt.Printf("sync start, goroutine=%d\n", runtime.NumGoroutine())
	af()
	fmt.Printf("sync end, goroutine=%d\n", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	fmt.Printf("async end, goroutine=%d\n", runtime.NumGoroutine())
}
