package main

import (
	"gee"
	"net/http"
)

//  curl -i http://localhost:9999/index
//HTTP/1.1 200 OK
//Content-Type: text/html
//Date: Tue, 13 Apr 2021 14:02:01 GMT
//Content-Length: 19
//
//<h1>Index Page</h1>%

// ❯ curl -i http://localhost:9999/indexx
//HTTP/1.1 404 Not Found
//Content-Type: text/plain
//Date: Tue, 13 Apr 2021 14:02:42 GMT
//Content-Length: 23
//
//404 NOT FOUND: /indexx

// ❯ curl -i http://localhost:9999/v1/
//HTTP/1.1 200 OK
//Content-Type: text/html
//Date: Tue, 13 Apr 2021 14:03:25 GMT
//Content-Length: 18
//
//<h1>Hello Gee</h1>%

// <h1>Hello Gee</h1>%                                                                                                                     ❯ curl "http://localhost:9999/v1/hello?name=geektutu"
//hello geektutu, you're at /v1/hello

func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")

}
