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
	// time.Stop() return false in this case because the timer already stopped because of the timer.C blocking
	// 2 seconds. so stop a stopped timer would give you false, otherwise you would get a true.
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}

/**
Present time: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001.
Expiration time: 2009-11-10 23:00:02 +0000 UTC m=+2.000000001.
Stop timer: false.
*/
