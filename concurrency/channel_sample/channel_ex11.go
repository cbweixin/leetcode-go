package channel_sample

import "fmt"

func Select_test2() {
	chanCap := 5
	intChan := make(chan int, chanCap)
	// each time , it would have a different output
	// so when mutiple case is satisifed, go would use some pesudo-random algo to choose a branch
	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}

/**
1
1
3
1
1
or
3
1
1
3
1
or
2
3
3
1
3
*/
