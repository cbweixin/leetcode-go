package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

// ❯ curl -i http://localhost:9999/
//HTTP/1.1 200 OK
//Content-Type: text/html
//Date: Fri, 16 Apr 2021 14:09:41 GMT
//Content-Length: 18
//
//<h1>Hello Gee</h1>%
//
// 2021/04/16 07:09:41 [200] / in 4.651µs

//  curl -i http://localhost:9999/v2/hello/cname
//HTTP/1.1 500 Internal Server Error
//Content-Type: application/json
//Date: Fri, 16 Apr 2021 14:13:30 GMT
//Content-Length: 36
//
//{"message":"Internal Server Error"}
//2021/04/16 07:13:30 [500] /v2/hello/cname in 114.09µs for group v2
//2021/04/16 07:13:30 [500] /v2/hello/cname in 142.122µs

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
