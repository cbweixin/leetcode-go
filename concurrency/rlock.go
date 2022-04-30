package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/**
	Try to lock for reading... [0]
	Locked for reading...[0]
	Try to lock for reading... [1]
	Locked for reading...[1]
	Try to lock for reading... [2]
	Locked for reading...[2]
	Try to lock for writing...
	Try to unlock for reading...[2]
	Unlocked for reading...[2]
	Try to unlock for reading...[1]
	Unlocked for reading...[1]
	Try to unlock for reading...[0]
	Unlocked for reading...[0]
	Locked for writing.

	*/
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			// read lock is not mutual exclusive
			fmt.Printf("Try to lock for reading... [%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked for reading...[%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to unlock for reading...[%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for reading...[%d]\n", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try to lock for writing...")
	// write lock
	rwm.Lock()
	fmt.Println("Locked for writing.")
}
