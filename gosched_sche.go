package main

import (
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		println("hello world")
	}()
	wg.Wait()
}
