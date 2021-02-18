package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	// the output is as below:
	//
	//  Hello, Mark!
	//	Hello, Mark!
	//	Hello, Mark!
	//	Hello, Robert!
	//	Hello, Mark!
	// name as an variable would be assigned to value 'Eric', .. 'Mark',because for loop is fast
	// also the time to run goroutine is not determinstic, so goroutine runs, very possibly the name
	// already has value Mark
	// if you want each name for each goroutine, see concur_go_2.go file
	for _, name := range names {
		go func() {
			fmt.Printf("Hello, %s!\n", name)
		}()
	}
	time.Sleep(time.Millisecond)

}
