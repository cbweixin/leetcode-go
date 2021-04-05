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
