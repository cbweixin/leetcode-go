package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

const (
	RateLimitPeriod = time.Second
	RateLimit       = 200
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
			fmt.Println("current token cnt:", len(quotas), time.Now())
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
func GetApi() {
	api := "http://localhost:8081/"
	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Printf("get api success\n")
	}
}

func Benchmark_Main(b *testing.B) {
	FillToken()
	fmt.Println(time.Now().Unix())
	//time.Sleep(time.Second)
	for i := 0; i < b.N; i++ {
		if TakeAvailable(true) {
			fmt.Println("get it")
		} else {
			fmt.Println("failed to get token")
		}
	}
}
