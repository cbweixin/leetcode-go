package main

import (
	"gee"
	"net/http"
)

// ❯ curl -i http://localhost:9999/
//HTTP/1.1 200 OK
//Content-Type: text/html
//Date: Mon, 05 Apr 2021 04:45:31 GMT
//Content-Length: 18
//
//<h1>Hello Gee</h1>%

//❯ curl -i http://localhost:9999/hello\?name\=xin
//HTTP/1.1 200 OK
//Content-Type: text/plain
//Date: Mon, 05 Apr 2021 04:46:38 GMT
//Content-Length: 28
//
//hello xin, you're at /hello

// ❯ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
//{"password":"1234","username":"geektutu"}

// ❯ curl -i http://localhost:9999/hellooo\?name\=xin
//HTTP/1.1 404 Not Found
//Content-Type: text/plain
//Date: Mon, 05 Apr 2021 04:48:11 GMT
//Content-Length: 24
//
//404 Not Found: /hellooo

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")

}
