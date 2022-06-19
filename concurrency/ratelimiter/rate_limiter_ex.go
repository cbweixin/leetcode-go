package ratelimiter

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func Serve() {
	r := rate.Every(1 * time.Millisecond)
	limiter := rate.NewLimiter(r, 10)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if limiter.Allow() {
			fmt.Printf("请求成功，当前时间：%s\n", time.Now().Format("2022-01-02 15:04:05"))
		} else {
			fmt.Printf("请求成功，但是被限流了。。。\n")
		}
	})

	_ = http.ListenAndServe(":8081", nil)
}
