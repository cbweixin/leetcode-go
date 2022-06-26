package main

import (
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer println("A.defer")

		func() {
			defer println("B.defer")
			runtime.Goexit()
			println("B")
		}()

		println("A")

	}()
}
