package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var count int
	var once sync.Once
	max := rand.Intn(100)

	for i := 0; i < max; i++ {
		// review check the implementation of `once.Do`, is very inspiring
		once.Do(func() {
			count++
		})
	}

	fmt.Printf("Count : %d.\n", count)

}
