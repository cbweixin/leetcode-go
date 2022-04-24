package channel_sample

import (
	"fmt"
	"time"
)

func Timer_test() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v.\n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}

/**
Present time: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001.
Expiration time: 2009-11-10 23:00:02 +0000 UTC m=+2.000000001.
Stop timer: false.
*/
