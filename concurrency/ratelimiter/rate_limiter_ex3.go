package ratelimiter

import (
	"time"
)

var quotas = make(chan time.Time, RateLimit)

func FillToken() {
	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()

		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()
}
func TakeAvailable(block bool) bool {
	var takenResult bool
	if block {
		select {
		case <-quotas:
			takenResult = true
		}
	} else {
		select {
		case <-quotas:
			takenResult = true
		default:
			takenResult = false
		}
	}
	return takenResult
}
