package ratelimiter

import (
	"fmt"
	"time"
)

type Request interface{}

func handle(r Request) {
	fmt.Println(r.(int))
}

const (
	RateLimitPeriod = time.Minute
	RateLimit       = 200
)

func HandleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

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

	for r := range requests {
		<-quotas
		go handle(r)
	}
}
